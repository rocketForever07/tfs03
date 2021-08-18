function isPalindrome(str){
    //remove remove non-alphanumeric elements
    let strFormat = str.replace(/[^a-zA-Z0-9]/g, '').toLowerCase();

    //compare to reverse string
    return strFormat==strFormat.split('').reverse().join('');

}

module.exports = isPalindrome 