function isPalindrome(str){
    //remove remove non-alphanumeric elements
    let strFormat = str.replace(/[^a-zA-Z0-9]/g, '').toLowerCase();

    return strFormat==strFormat.split("").reverse().join("");

}

module.exports = isPalindrome 