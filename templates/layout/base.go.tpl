{{define "base"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>

    <script
      src="https://kit.fontawesome.com/05ce6c0c26.js"
      crossorigin="anonymous"
    ></script>
    
    <link rel="stylesheet" href="/static/css/main.css" />
  </head>
  
  <body>
    {{template "body" .}}
    <script src="/static/scripts/charts.js"></script>
  </body>
</html>
{{ end }}
