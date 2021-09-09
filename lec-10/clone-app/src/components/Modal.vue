<template>
  <div class="container-modal" :class="{ showModal: showModal}">
      <div class="modal">
          <div class="header">
              <button @click="closeButton">x</button>
          </div>
          <div class="body-modal">
              <div class="main-body">
                  <h4>Success!</h4>
                  <img src="../assets/images/success.png">
              </div>
              <div class="foot-body" @click="closeButton">
                  <btn>OK</btn>
              </div>
          </div>
      </div>
  </div>

</template>

<script>
export default {
    name:"Modal",
    props:['showModal'],

    methods: {
        closeButton(){
            this.$emit('closeButton',true)
        }
    }
}

</script>

<style scope>
    .showModal{
        display: none;
    }
    .container-modal{
        width: 100%;
        height: 100%;
        background:whitesmoke;
        position: relative;
        
    }
        .modal{
            width:500px;
            height:300px;
            position:absolute;
            top:calc(50% - 150px);
            left: calc(50% - 250px);
            border-radius: 5px;
        }
            .header{
                width: 100%;
                height: 10%;
                background:royalblue;
                border-radius: 5px 5px 0px 0px;

            }
                .header button{
                    float:right;
                    height: 100%;
                    width: 8%;
                    background: red;
                    border: none;
                    cursor: pointer;
                    border-radius: 0px 5px 0px 0px;

                }
            .body-modal{
                border-radius: 0px 0px 5px 5px;

                width: 100%;
                height: 90%;
                background: white;
                padding-top: 20px;

            } 
                img{
                    width: 25%;
                    height:50%;

                }   
            .foot-body{
                margin: auto;
                width: 10%;
                size: 10%;
                line-height: 200%;
                background: royalblue;
                color: aliceblue;
                border-radius: 2px;
                cursor: pointer;
            }    

</style>