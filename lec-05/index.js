function isPalindrome(str){
    //remove remove non-alphanumeric elements
    let strFormat = str.replace(/[^a-zA-Z0-9]/g, '').toLowerCase();

    return strFormat==strFormat.split("").reverse().join("");

}

function  uniqueUnion(...arrs){
    //return a set
    return [...new Set([...arrs])];
}

function seekAndDestroy(arr,...arg){
    //convert to list
    destroyList=[...arg];

    destroyList.forEach(element => {
        const index = arr.indexOf(element);
        if (index > -1) {
            //use splice to remove
            arr.splice(index, 1);
        }
    });

    return arr;
};

function toSpinalCase(str) {

    let newStr = str.replace(/([a-z])([A-Z])/g, ('$1 $2'));
    //replace space with '-'
    newStr = newStr.replace(/ /g, '-');
    //replace operator with '-'
    newStr = newStr.replace(/_/g,'-');


    return newStr.toLowerCase()
};



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


