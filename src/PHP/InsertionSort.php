<?php

/*
*Function for sorting an array with insertion sort algorithms
*/


function sortInsertion($array) {
    $sortedArray = array();
    for ($i = 0 ; $i < count($array); $i++) {
        $element = $array[$i];
        $j = $i;
        while($j > 0 && $sortedArray[$j-1] > $element) {
           $sortedArray[$j] = $sortedArray[$j-1];
            $j = $j-1;
        }
        $sortedArray[$j] = $element;
    }
    return $sortedArray;
}

$unsorted = [5,2,6,8,4];
$val = sortInsertion($unsorted);
print_r($val);

?>
