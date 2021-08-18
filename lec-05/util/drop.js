function drop(arr,func){
    for(let i=0;i<arr.length;i++){
        if(func(arr[i])==false){
            arr.splice(i, 1);
            i--;
        }else{
            break;     
        }
    }

    return arr;
}

module.exports = drop