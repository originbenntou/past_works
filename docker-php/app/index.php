<?php

require_once(__DIR__ . '/config.php');
require_once(__DIR__ . '/TaskProvider.php');

$model = new TaskProvider();
var_dump($model);

?>
<html lang="ja">
  <head>
    <title>Task Manager</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="../css/bootstrap.min.css" type="text/css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>

    <!-- GoogleFusionTable連携
    <script>
      var apiKey = 'AIzaSyBer987AbVXVOYb4di0IswTNEZsQUpcxEs';
      var tableId = '1ZVB9MT_Pf9I-ERq0FPbU_N_x--V7HxNXzL01AJHY';

      function initialize(){
        var sql = 'SELECT * FROM ' + tableId;
        var url = 'https://www.googleapis.com/fusiontables/v1/query';
        url += '?key=' + apiKey;
        url += '&sql=' + encodeURIComponent(sql);

        $.ajax({
          url: url,
          dataType: 'json',
          success: function (data) {
            console.log(data);
          }
        });
      }
      initialize();
    </script>
    -->
  </head>
  <body>
    <section>
      <h1>Task Manager</h1>
      <div id="undone" class="container-fluid">
        <h2>未処理</h2>
        <div class="form-group">
          <input class="form-control" type="text" placeholder="What needs to be done?">
        </div>
        <table class="table table-striped">
          <tr><td>やることやることやることやることやることやることやることやること</td><td>いつまで</td></tr>
          <tr><td>やること</td><td>いつまで</td></tr>
          <tr><td>やること</td><td>いつまで</td></tr>
          <tr><td>やること</td><td>いつまで</td></tr>
        </table>
      </div>
      <div id="doing" class="container-fluid">
        <h2>処理中</h2>
        <div class="form-group">
          <input class="form-control" type="text" placeholder="What needs to be done?">
        </div>
        <table class="table table-striped">
          <tr><td>やることやることやることやることやることやることやることやること</td><td>いつまで</td></tr>
          <tr><td>やること</td><td>いつまで</td></tr>
          <tr><td>やること</td><td>いつまで</td></tr>
          <tr><td>やること</td><td>いつまで</td></tr>
        </table>
      </div>
    </section>
  </body>
</html>
