# Docker for Laravel
Docker + Laravel + MySQL で環境構築する

## 参考サイト
### DockerでLaravel環境構築
[Dockerでサクっとローカル開発環境(LAMP + Laravel)を構築する](https://qiita.com/ProjectEuropa/items/5fbb00848cd8d5b57182)

### Laravelで作成したDBにCRUDできるかどうか(Uはしてないｗ)
[Laravel入門: 初心者でも10分でWebサービスを作れる！PHPフレームワークLaravelとPaizaCloudの使い方 - paiza開発日誌](https://paiza.hatenablog.com/entry/2018/02/16/paizacloud_laravel)

## ゴール
- http://localhost/ でLaravelのデフォルト画面を表示させる
- DBを作成してアクセスできるかどうか確認する

## 手順
### ディレクトリ構成
docker-laravel/
├── apache-php/
│　　├── Dockerfile
│　　└── apache.conf
├── lara-d
└── docker-compose.yml

### 1. Laravelをインストールする
```
composer create-project "laravel/laravel=5.5.*" lara-d
```

composer入ってない場合は入れる
[Macにhomebrewでcomposerをインストール](https://qiita.com/yamatmoo/items/b6d234f33929f07c43d4)

### 2. docker-compose.ymlとDockerfileを作成
#### Dockerfile
```
FROM centos:7.4.1708

RUN yum -y update

# 外部リポジトリ（EPEL, Remi）を追加
RUN yum -y install epel-release
RUN yum -y install http://rpms.remirepo.net/enterprise/remi-release-7.rpm

# apache その他 phpパッケージを導入
RUN yum -y install httpd
RUN yum -y install --enablerepo=remi,remi-php72 php php-cli php-common php-devel php-fpm php-gd php-mbstring php-mysqlnd php-pdo php-pear php-pecl-apcu php-soap php-xml php-xmlrpc 
RUN yum -y install zip unzip

# composerのインストール
RUN curl -sS https://getcomposer.org/installer | php
RUN mv composer.phar /usr/local/bin/composer

CMD ["/usr/sbin/httpd","-D","FOREGROUND"]

WORKDIR /var/www/html
```

#### docker-compose.yml
```
version: "3"
services:
  web:
    build:
      context: ./apache-php
    ports: 
      - 80:80
    privileged: true
    links:
      - db
    volumes:
      - "./lara-d/:/var/www/html"
      - "./apache-php/apache.conf:/etc/httpd/conf/httpd.conf"
    container_name: "apache-php"
  db:
    image: mysql:5.7
    environment:
      - MYSQL_DATABASE=testdb
      - MYSQL_ROOT_PASSWORD=root
    ports:
      - 3306:3306
    container_name: "mysql"

```

#### httpd.confの設定（をapache.confに書いてマウントさせる）
```
# ホスト側httpd.confをコピペして修正
## 119
- DocumentRoot "/var/www/html"
+ DocumentRoot "/var/www/html/public"

## 124
- <Directory "/var/www">
+ <Directory "/var/www/html/public">

## 125
- AllowOverride None
+ AllowOverride All

## 追加
LoadModule rewrite_module modules/mod_rewrite.so
```

### 3. buildを実行してイメージを作成する
```
$ sudo docker-compose build
```

### 4. コンテナ起動
```
$ sudo docker-compose up -d
```

### 5. Laravelデフォルト画面表示
http://localhost/

### 6. DB接続確認
```
$ mysql -u root -h 0.0.0.0 -port 3306 -p
```

一旦これでOK

## Laravelで軽くMVCしてみる
### DB接続設定
docker-laravel/lara-d/.env
```
# 9
- DB_HOST=127.0.0.1
+ DB_HOST=mysql5.7

# 11
- DB_DATABASE=homestead
+ DB_DATABASE=testdb

# 12
- DB_USERNAME=homestead
+ DB_USERNAME=root

# 13
- DB_PASSWORD=secret
+ DB_PASSWORD=root
```

### モデル・コントローラを作成
`artisan make:model`コマンドでモデル、コントローラ、DBマイグレーションファイルができる
```
$ docker exec apache-php php artisan make:model Task -m -c -r
```

### データベース作成
`database/migrations/2018_xx_xx_xxxxxxxx_create_tasks_table`という形式でマイグレーションファイルができているので、`$table->string('name');`を加える
```
    public function up()
    {
        Schema::create('tasks', function (Blueprint $table) {
            $table->increments('id');
            $table->string('name');
            $table->timestamps();
        });
    }
```

実行
```
$ docker exec apache-php php artisan migrate
```

確認
```
# nameカラムがある
originbenntou:docker-laravel[master *+]$ docker exec -it mysql mysql -u root -p testdb -e "desc tasks"
Enter password: 
+------------+------------------+------+-----+---------+----------------+
| Field      | Type             | Null | Key | Default | Extra          |
+------------+------------------+------+-----+---------+----------------+
| id         | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| name       | varchar(255)     | NO   |     | NULL    |                |
| created_at | timestamp        | YES  |     | NULL    |                |
| updated_at | timestamp        | YES  |     | NULL    |                |
+------------+------------------+------+-----+---------+----------------+

```

### ルーティング設定
routes/web.php
```
Route::get('/tasks', 'TaskController@index');
Route::post('/tasks', 'TaskController@store');
Route::delete('/tasks/{id}', 'TaskController@destroy');
```

### コントローラ設定
app/Http/Controllers/TestController.php
```
<?php

namespace App\Http\Controllers;

use App\Task;
use Illuminate\Http\Request;

class TaskController extends Controller
{
    public function index()
    {
        $tasks = Task::all();
        return view('tasks', ['tasks' => $tasks]);
    }

    public function store(Request $request)
    {
        $task = new Task;
        
        $task->name = request('name');
        $task->save();
        return redirect('/tasks');
    }

    public function destroy(Request $request, $id, Task $task)
    {
        $task = Task::find($id);
        $task->delete();
        return redirect('/tasks'); 
    }
}
```

### View設定
resources/views/task.blade.php
```
<!DOCTYPE html>
<html>
<head>
    <title>Task List</title>
    <!-- CSS And JavaScript -->
    <link rel="stylesheet" href="//fonts.googleapis.com/css?family=Roboto:300,300italic,700,700italic">
    <link rel="stylesheet" href="//cdn.rawgit.com/necolas/normalize.css/master/normalize.css">
    <link rel="stylesheet" href="//cdn.rawgit.com/milligram/milligram/master/dist/milligram.min.css">
</head>
<body>
<div class="container">
    <h1>Task List</h1>
    <form action="/tasks" method="POST" class="form-horizontal">
        {{ csrf_field() }}
        <!-- Task Name -->
        <div class="form-group">
            <label for="task" class="col-sm-3 control-label">Task</label>
            <div class="col-sm-6">
                <input type="text" name="name" id="task-name" class="form-control">
            </div>
        </div>

        <!-- Add Task Button -->
        <div class="form-group">
            <div class="col-sm-offset-3 col-sm-6">
                <button type="submit" class="btn btn-default">
                    <i class="fa fa-plus"></i> Add Task
                </button>
            </div>
        </div>
    </form>

    <!-- Current Tasks -->
    <h2>Current Tasks</h2>
    <table class="table table-striped task-table">
        <thead>
        <th>Task</th><th>&nbsp;</th>
        </thead>

        <tbody>
        @foreach ($tasks as $task)
        <tr>
            <!-- Task Name -->
            <td>
                <div>{{ $task->name }}</div>
            </td>
            <td>
                <form action="/tasks/{{ $task->id }}" method="POST">
                    {{ csrf_field() }}
                    {{ method_field('DELETE') }}
                    <button>Delete Task</button>
                </form>
            </td>
        </tr>
        @endforeach
        </tbody>
    </table>
</div>
</body>
</html>
```

### 画面で見てみる
http://localhost/test
