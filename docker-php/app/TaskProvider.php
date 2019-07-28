<?php

require_once(__DIR__ . '/DbClient.php');

class TaskProvider extends DbClient
{
    public function getUndone() {
        $model = (new DbClient)->getAll();
        return array_filter ($model, function ($e) {
            return $e->status == 0;
        });
    }
}