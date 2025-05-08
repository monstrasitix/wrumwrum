{{define "table"}}

{{$headings := .Headings}}

<div class="table">
  <table>
    <thead>
      <tr>
        {{range $heading := $headings}}
          <th style="min-width: {{$heading.Width}}px" class="{{tableAlign $heading.Type}}">
            {{$heading.Text}}
          </th>
        {{end}}
      </tr>
    </thead>

    <tbody>
      {{range $row := .Rows}}
        <tr>
          {{range $i, $value := $row}}
            <td class="{{tableAlign (index $headings $i).Type}}">
              {{$value}}
            </td>
          {{end}}
        </tr>
      {{end}}
    </tbody>
  </table>
</div>
{{end}}
