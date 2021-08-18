function toSpinalCase(str) {

    let newStr = str.replace(/([a-z])([A-Z])/g, ('$1 $2'));
    //replace space with '-'
    newStr = newStr.replace(/ /g, '-');
    //replace operator with '-'
    newStr = newStr.replace(/[_*+//\\]/g,'-');


    return newStr.toLowerCase()
};

module.exports = toSpinalCase