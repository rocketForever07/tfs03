function  uniqueUnion(...args){

    //convert args to array
    arr = [].concat(...args);

    //return a set
    return [...new Set([...arr])];
}

module.exports=uniqueUnion