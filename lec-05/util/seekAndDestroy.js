
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

module.exports = seekAndDestroy
