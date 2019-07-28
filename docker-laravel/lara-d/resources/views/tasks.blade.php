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
