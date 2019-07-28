<?php

require_once(__DIR__ . '/config.php');

class DbClient
{
    private $_db;

    public function __construct() {
        try {
            $this->_db = new PDO(DB_HOST, DB_USER, DB_PASSWORD);
            $this->_db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        } catch (PDOException $e) {
            echo $e->getMessage();
        }
    }

    public function getAll() {
        $model = $this->_db->query("select * from tasks order by id desc");
        return $model->fetchAll(PDO::FETCH_OBJ);
    }
}