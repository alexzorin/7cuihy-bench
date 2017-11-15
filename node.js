const Benchmark = require('benchmark')
const suite = new Benchmark.Suite;

var generateRandomArray = function (size) {
  var myArray = [];
  myArray.length = size
  for (var i = 0; i < myArray.length; i++){
    myArray[i] = Math.floor(Math.random() * 1000) + 1;
  }
  return myArray;
};

const doTheThing = function (data, greater, lesser) {
  let numGreater = 0
  let numLesser = 0
  for (var i = 0; i < data.length; i++){
    let v = data[i]
    if (v > 500){
      greater[numGreater++] = v
    } else if(v < 500) { 
      lesser[numLesser++] = v
    }
  }
  // Truncate lesser/greater to their actual sizes
  // lesser.length = numLesser
  // greater.length = numGreater
}

// Pre-allocate all of these
// to avoid extra heap allocation during calculation
const data = generateRandomArray(1e6) // 1mil elements
const lesser = []
lesser.length = data.length
const greater = []
greater.length = data.length

suite
.add('Sort the data', function () {
  doTheThing(data, greater, lesser)
})
.on('cycle', function (evt) {
  console.log(evt.target.toString())
})
.on('error', function (e) {
  console.log(e)
})
.run()
