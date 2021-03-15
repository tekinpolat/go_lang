<?php 


function testarr($arr){
    $arr[0] = 21;
}

function testobject($obj){
    $obj->a = 43;
}
$arr = [1,2,3,4];
//call by value
testarr($arr);

print_r($arr);

$arr = ['a'=>21, 'b'=>43];
$val = (object)$arr;
//call by reference
testobject($val);
print_r($val);


#mutable 
#immutable

$name = "Tekin";
$name[0] = 'M';

echo $name.PHP_EOL;

