<?php
var_dump($_GET);
$password = $_GET["password"];
$returnValue=[];
$returnValue["success"]=false;
if ($password==123 && $_GET["username"]=="james")
	$returnValue["success"]=true;
echo json_encode($returnValue);
?>
