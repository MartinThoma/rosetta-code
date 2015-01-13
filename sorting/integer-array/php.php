<?php

$n = 10000000;
$max = 1000000000;
$array = [];

for ($i=0; $i < $n; $i++) { 
    $array[] = rand(0, $max);
}

$t0 = microtime(TRUE);
sort($array);
$t1 = microtime(TRUE);

printf("%0.2f seconds\n", $t1-$t0);

?>