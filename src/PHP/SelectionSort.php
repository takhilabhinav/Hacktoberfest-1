<?php

function selectionSort(array $arr){

$n =sizeof($arr);
    for($i = 0; $i < $n; $i++) {
        $lowestValueIndex = $i;
        $lowestValue = $arr[$i];
        for($j = $i+1; $j < $n; $j++) {
            if($arr[$j] < $lowestValue) {
                $lowestValueIndex = $j;
                $lowestValue = $arr[$j];
            }
        }
        $arr[$lowestValueIndex] = $arr[$i];
        $arr[$i] = $lowestValue;
   }
  return $arr;
}
// Example:
$arr = array(64,25,12,22,11);
$result = selectionSort($arr);
print_r($result);

?>