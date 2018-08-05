
var Year = function (year) {
    this.value = year;
};

Year.prototype.isLeap = function () {

    if(this.isDivisible(this.value,400)){
        return true;
    }

    if(this.isDivisible(this.value,100)){
        return false;
    }
    if(this.isDivisible(this.value,4)){
        return true;
    }

    return false;
};

Year.prototype.isDivisible = function(dividend, divisor){
  if (divisor == 0) {
      return false;
  } 
  return (dividend % divisor == 0)
};
module.exports = Year;
