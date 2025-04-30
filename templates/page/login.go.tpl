{{template "base" .}}

{{define "body"}}
<div class="wallpaper --vh-100">
  <img
    class="wallpaper-image"
    src="/static/assets/images/aston-martin.jpeg"
    alt="Background"
  />

  <div class="wallpaper-content">
    <div class="flex -vertical -align-center -justify-center --vh-100">
      <div class="container -modal">
        <div class="card" style="background-color: rgba(255, 255, 255, 0.2)">
          <div class="card-body">
            {{template "login-form" .}}
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
{{ end }}

{{define "login-form"}}
<form method="POST" class="form" novalidate autocomplete="off">
  <div class="form-control">
    <label class="label" style="color: white" for="email">Email</label>
    <input class="input" id="email" name="email" type="email" />
  </div>

  <div class="form-control">
    <label class="label" style="color: white" for="password">Password</label>
    <input class="input" id="password" name="password" type="password" />
  </div>

  <div class="form-control">
    <button type="submit" class="button -primary">Login</button>
  </div>
</form>
{{ end }}
