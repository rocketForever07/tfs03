<template>
  <div class="container">
        <input type="text" id="result" :value="displayText">
        <table>
            <tr>
                <td><input type="button" value="1" @click="mainButton('1')" class=" display"></td>
                <td><input type="button" value="2" @click="mainButton('2')" class=" display"></td>
                <td><input type="button" value="3" @click="mainButton('3')" class=" display"></td>
                <td><input type="button" value="+" @click="mainButton('+')" class=" display"></td>
            </tr>
            <tr>
                <td><input type="button" value="4" @click="mainButton('4')" class=" display"></td>
                <td><input type="button" value="5" @click="mainButton('5')" class=" display"></td>
                <td><input type="button" value="6" @click="mainButton('6')" class=" display"></td>
                <td><input type="button" value="-" @click="mainButton('-')" class=" display"></td>
            </tr>
            <tr>
                <td><input type="button" value="7" @click="mainButton('7')" class=" display"></td>
                <td><input type="button" value="8" @click="mainButton('8')" class=" display"></td>
                <td><input type="button" value="9" @click="mainButton('9')" class=" display"></td>
                <td><input type="button" value="x" @click="mainButton('x')" class=" display"></td>
            </tr>
            <tr>
                <td><input type="button" value="0" @click="mainButton('0')" class=" display"></td>
                <td><input type="button" value="Del" @click="mainButton('Del')" id="del"></td>
                <td><input type="button" value="%" @click="mainButton('%')" class=" display"></td>
                <td><input type="button" value="/" @click="mainButton('/')" class=" display"></td>
            </tr>
            <tr>
                <td colspan="4"><input type="button" value="=" @click="mainButton('=')" id="exc"></td>
            </tr>
        </table>
    </div>
</template>

<script>
import * as axios from "axios";


export default {
  name: 'App',
  data(){
    return {
        displayText:"",
        displayResult: false,
        // mainButtonList: ['1' ,'2','3','+','4','5','6','-','7','8','9','x','0','%','/','Del','='],
    };
  },

  methods: {
      mainButton(ch){
          if(ch==="Del"){
              this.displayText="";
          }else if(ch==="="){
              this.excute(this.displayText)
              this.displayResult=true;
          }else{
              if(this.displayResult==true){
                this.displayResult=false;
                this.displayText="";
              }
              this.displayText+=ch;
          }
      },

      excute(expression){
        console.log(expression)
        this.displayText="ok";

        var seft=this;

        axios.post("localhost:10000/calculator",expression)
        .then(function(res){
          seft.displayText=res.data;
        })
        .catch(function(res){
          seft.displayText="call api";
          console.log(res)
        })
      }
  },

}
</script>

<style scoped>
    .container{
        position: absolute;
        top: 100px;
        left: calc(50% - 90px);
    }
    table{
        width: 150px;
    }
    #result{
        width: 180px;
        font-size: 15px;
        padding: 16px 5px;
        text-align: right;
        border: none;
        background-color: #F8F8F8;
    }
    input{
        padding: 15px;
        font-weight: bold;
        font-size: 16px;
        border: none;
        border-radius: 3px;
        background-color: #F2F2F2;
        width: 100%;
    }
    input:hover{
        background-color: rgb(171, 175, 162);
    }

    #exc{
        background-color: rgb(75, 223, 107);
    }
    #exc:hover{
        background-color: #12cf3b;
    }

    #del{
        background-color: tomato;

    }
    #del:hover{
        background-color: rgb(197, 43, 16);
    }
</style>>
