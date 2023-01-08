## get-grocery-data

El siguiente repositorio contiene código que se utiliza para scrapear información de productos de un sitio web (EJ: Supermercado).<br />
Esta información se recopila mediante la realización de una solicitud HTTP a una página web y el análisis del contenido HTML de la respuesta utilizando la biblioteca goquery.
<br /><br />
Adicionalmente, la información obtenida es exportada como .csv hacia una ruta específica en el equipo y hacia una base de datos postgres.<br />

![Grocerydata](https://user-images.githubusercontent.com/64023919/211224999-f57e1ae9-ce0f-4af1-9f93-910b537768b5.png)

- Importación:<br />
``
go get github.com/AlexiaVtr/get-grocery-data
``
<br />

Modifica "scraping.go" para adaptarlo a la página que que quieras hacer scraping.<br />
