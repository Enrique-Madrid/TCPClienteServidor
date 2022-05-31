<!doctype html>
<html>
    <head>
    <link rel="shortcut icon" href="#" />
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    
    <!-- Bootstrap CSS -->    
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
    <!-- FontAwesom CSS -->
    <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.8.1/css/all.css" integrity="sha384-50oBUHEmvpQ+1lW4y57PTFmhCaXp0ML5d60M1M7uH2+nqUivzIebhndOJK28anvf" crossorigin="anonymous">            
    <!--CSS custom -->  
    <link rel="stylesheet" href="main.css">  
    </head>
    <body>
    <header class="text-center">
        <h2 class="d-inline p-2 bg-primary text-white"><span class="badge badge-success">CRUD con VUE.JS</span></h2>
    </header>    
    
     <div id="appDatos">               
        <div class="container">                
            <div class="row">
                <div class="col text-right">                        
                    <h5>Stock Total: <span class="badge badge-success">{{totalPeso}}</span></h5>
                </div>    
            </div>                
            <div class="row mt-5">
                <div class="col-lg-12">                    
                    <table class="table table-striped">
                        <thead>
                            <tr class="bg-primary text-light">
                                <th>ID</th>                                    
                                <th>Marca</th>
                                <th>Modelo</th>
                                <th>Stock</th>    
                            </tr>    
                        </thead>
                        <tbody>
                            <tr v-for="(datos,indice) of datos">                                
                                <td>{{datos.id}}</td>                                
                                <td>{{datos.canal}}</td>
                                <td>{{datos.nombre}}</td>
                                <td>
                                    <div class="col-md-8">
                                    <input type="number" v-model.number="datos.peso" class="form-control text-right" disabled>      
                                    </div>    
                                </td>
                            </tr>    
                        </tbody>
                    </table>                    
                </div>
            </div>
        </div>        
    </div>        
      
    <!--Vue.JS -->    
    <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>              
    <!--Axios -->      
    <script src="https://cdnjs.cloudflare.com/ajax/libs/axios/0.15.2/axios.js"></script>     
    <!--Código custom -->          
    <script src="main.js"></script>         
    </body>
</html>