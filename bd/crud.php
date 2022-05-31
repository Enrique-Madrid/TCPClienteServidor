<?php
include_once 'conexion.php';
$objeto = new Conexion();
$conexion = $objeto->Conectar();

$_POST = json_decode(file_get_contents("php://input"), true);
$opcion = (isset($_POST['opcion'])) ? $_POST['opcion'] : '';
$id = (isset($_POST['id'])) ? $_POST['id'] : '';
$canal = (isset($_POST['canal'])) ? $_POST['canal'] : '';
$nombre = (isset($_POST['nombre'])) ? $_POST['nombre'] : '';
$peso = (isset($_POST['peso'])) ? $_POST['peso'] : '';

$consulta = "SELECT id, canal, nombre, peso FROM datosenviados2";
$resultado = $conexion->prepare($consulta);
$resultado->execute();
$data=$resultado->fetchAll(PDO::FETCH_ASSOC);

print json_encode($data, JSON_UNESCAPED_UNICODE);
$conexion = NULL;