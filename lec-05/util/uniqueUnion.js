function  uniqueUnion(...arrs){
    //return a set
    return [...new Set([...arrs])];
}

module.exports=uniqueUnion