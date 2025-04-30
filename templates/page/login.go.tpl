{{template "base" .}}

{{define "body"}}
<div class="flex --vh-100">
  <div class="grow" style="background-color: dodgerblue">
    <h1></h1>
  </div>

  <div class="grow">
    <div class="container -modal --h-100">
      <div class="flex -vertical -align-center -justify-center --h-100">
        <form method="POST" class="form" novalidate autocomplete="off">
          <div class="form-control">
            <label class="label" for="email">Email</label>
            <input class="input" id="email" name="email" type="email" />
          </div>

          <div class="form-control">
            <label class="label" for="password">Password</label>
            <input
              class="input"
              id="password"
              name="password"
              type="password"
            />
          </div>

          <div class="form-control">
            <button type="submit" class="button -primary">Login</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</div>
{{ end }}
