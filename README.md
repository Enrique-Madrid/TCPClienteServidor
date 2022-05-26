Programa cliente / servidor

Este programa sirve para enviar archivos a traves de un servidor hacia un cliente

Para inicializar el cliente y escuchar, se usa el comando ./client log -canal-
Luego de eso, se selecciona el canal al cuál enviar el archivo, con el comando ./client channel -canal-
Luego de haber elejido un cliente para escuchar y otro para enviar, con el comando ./client send -archivo- puede enviar un archivo, y este se enviará al canal seleccionado anteriormente.