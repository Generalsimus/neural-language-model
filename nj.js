class Perceptron{constructor(){
    //bias , input1, input2
    this.inputs = [1,0,0];
    this.inputWeights = [(Math.random()*2)-1,(Math.random()*2)-1,(Math.random()*2)-1];
    this.output = 0;
    this.desiredOutput = 0; 
}//perceptron methods
    activate = () => {
        let sum = 0;
        for(var n = 0; n < this.inputs.length; n++){
            sum += this.inputs[n] * this.inputWeights[n];
        };
        this.output = sum < 0 ? 0 : 1;
        this.desiredOutput == this.output ? console.log("Correct answer") : console.log("Incorrect answer");
    };
    propagate = () => {
        let error = this.desiredOutput - this.output;
        for(var m = 0; m < this.inputs.length; m++){
            let delta = error * this.inputs[m];
            this.inputWeights[m] = this.inputWeights[m] + (delta * learningRate);
        } 
    };
}
let learningRate = 0.1;
var train = (iterations) => {
    for(var x = 0; x < iterations; x++){
        for(var y = 0; y < dataset.length; y++){
        perception.inputs = [1,dataset[y][0],dataset[y][1]];
        perception.desiredOutput = dataset[y][2];   
        perception.activate();
        perception.propagate(); 
        }
    }
}   
var perception = new Perceptron();
//[input1 , input2 , desiredOutput] 
var dataset = [
    [0,0,1],
    [1,1,0],
    [0.1,0.3,1],
    [1.5,1.8,0]
];  
train(200); 
console.log(perception)