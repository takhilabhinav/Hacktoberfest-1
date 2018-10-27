function selection_Sort(a)
{

    var n = a.length;
    var x = a;

    for (var i = 0; i < n-1; i++){
        var min_index = i;
        for (var j = i+1; j < n; j++){
            if (x[j] < x[min_index]){
                min_index = j;
            }
        }
    }

    var temp = x[min_index];
    x[min_index] = x[i];
    x[i] = temp;

    return x;
}
