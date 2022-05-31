var url = "bd/crud.php";

var appDatos = new Vue({    
el: "#appDatos",   
data:{     
    datos:[],          
     canal:"",
     nombre:"",
     peso:"",
     total:0,       
 },    
methods:{  
    //PROCEDIMIENTOS para el CRUD     
    listarDatos:function(){
        axios.post(url).then(response =>{
           this.datos = response.data;       
        });
    },     
},      
created: function(){            
   this.listarDatos();            
},    
computed:{
    totalPeso(){
        this.total = 0;
        for(dato of this.datos){
            this.total = this.total + parseInt(dato.peso);
        }
        return this.total;   
    }
}    
});