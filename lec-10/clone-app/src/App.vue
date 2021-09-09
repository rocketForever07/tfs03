<template>
  <div class="container" >
    <div class="body" :class="{isHide: !showModal}">
      <div class="banner"></div>

      <div class="main-body">

        <div class="body-content">

          <div class="intro">
            <h4>Order Summary</h4>
            <p>You can now listen to milions of songs,</p>
            <p>audibooks,and podcasts on any device</p>
            <p>any where you like</p>
          </div>

          <div class="infor-packet">
            <div class="left-packet">   
              <div class="icon-music"></div>
              <div class="detail-packet">
                <b>{{selectedItem.name}}</b>
                <p>{{selectedItem.cost}}</p>
              </div>
            </div> 
           <div class="right-packet">
              <a @click="changePkg"><u>Change</u></a>
            </div>
          </div>
          <div class="process">
            <button class="btn-payment" @click="processBtn">Process to payment</button>
            <button class="btn-cancel">Cancel</button>
          </div>
      </div>
      </div>
    </div>

    <Modal :showModal="showModal" @closeButton="closeButton"/>
    
  </div>

</template>
<script>

import Modal from './components/Modal.vue';

export default {
  name: 'App',
  components: {
    Modal,
  },
  data(){
    return {
      msg:"khoi",
      first:"khoi",
      last:"doan",
      showModal:true,
      packetItems:[
        {id:1,name:"Anual Plane",cost:"$55,99/year"},{id:2,name:"Pro Plane",cost:"$99,99/year"}
      ],
      selectedItem:{},

    };
  },

  methods: {
    greet(event){
      alert("hello "+this.first)

      if(event){
        alert(event.target.tagName)
      }
    },

    changePkg(){
      if(this.selectedItem.id==1){
        this.selectedItem=this.packetItems[1];//pkg 2
      }else{
        this.selectedItem=this.packetItems[0];//pkg 1
      }
    },
    closeButton(value){
      this.showModal=value;
    },
    processBtn(){
      this.showModal=false;

      console.log(this.showModal);
      
    }
  },

  computed: {
    msgComputed(){
      return this.msg+"0711"
    },
    fullname(){
      return this.first+" "+this.last;
    }
  },

  watch: {
    msg(_,oldMsg){
      console.log(" ",oldMsg)
    },
    msgComputed(){
      console.log("computed change")
    }
  },

  created(){
    this.selectedItem=this.packetItems[0];
  }

}
</script>

<style scope>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.isHide{
  display: none;
}

.container{
  width: 100%;
  height: 550px;
  background-image: url("./assets/images/pattern-background-desktop.svg");
  background-repeat: no-repeat;
  background-position: 0px 0px;
  
}


.body{
  margin: auto;
  margin-top:30px;
  width: 20%;
  height: 100%;
  border-radius: 10px;
  
}

.banner{
  width: 100%;
  height: 30%;
  background-image: url("./assets/images/illustration-hero.svg");
  background-repeat: no-repeat;
  background-position: center;
  border-radius: 10px 10px 0px 0px;
}

.main-body{
  background:#ffffff;
  margin-top:0px  !important;
  width: 100%;
  height: 70%;
  font-size: 10px;
  border-radius: 0px 0px 10px 10px;
  
}

  .body-content{
    width: 80%;
    height: 100%;
    margin: auto;
  }
    .intro{
      width: 100%;
      height: 35%;
    }

      h4{
        margin-bottom: 17px;
        font-size: 2em;
        margin-top: 0px;
      }
      p{
          height: 5px;
          font-size: 1.2em;
      }
    .infor-packet{
      width: 100%;
      height: 25%;
      display: flex;
      border-radius: 12px;
      background: rgb(248 249 254);
      margin-bottom: 5px;
      margin-top: 0px;  
    }
      .left-packet{
        width: 70%;
        height: 100%;
        display: flex;
      }
        .icon-music{
          width: 40%;
          height: 100%;
          background-image: url("./assets/images/icon-music.svg");
          background-repeat: no-repeat;
          background-position: center;
        }
        .detail-packet{
          width: 60%;
          height: 100%;
          padding-top:20px;
        }
          b{
            margin-bottom: 0px;
            font-size: 1.2em;

          }
      .right-packet{
        width: 30%;
        height: 100%;
        padding-top:40px;

      }
        a{
          margin-top: 10px;
          color: blue;
          cursor: pointer;
        }

    .process{
      width: 100%;
      height: 40%;
      display: block;
    }    
      .btn-payment{
        display: block;
        margin: auto;
        width: 100%;
        height: 30%;
        border-radius: 5px;
        background:rgb(56 42 225);
        color: white;
        cursor: pointer;
        border: none;
      }
      .btn-cancel{
        margin: auto;
        margin-top:10px ;
        width: 100%;
        background: white;
        border: none;
      }


</style>
