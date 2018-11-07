</br>
</br>

<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>All Appointments</h1>
     <p class="lead">Here you'll find all your appointments</p>
     {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

    {{if .flash.notice}}
    <blockquote>{{.flash.notice}}</blockquote>
    {{end}}
   </div>
 </div>
</div>
</div>
<div class="container">
  <div class="row">
    <table class="table table-striped">
      <thead>
        <tr>
          <th>Id</th>
          <th>Title</th>
          <th>Date</th>
          <th>Time</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .records}}
        <tr>
          <td>{{$record.Id}}</td>
          <td>{{$record.Title}}</td>
          <td>{{$record.Date}}</td>
          <td>{{$record.Time}}</td>
          <td>
            <a href="/appointment/edit/{{$record.Id}}" class="btn btn-primary">Edit</a>
            <a href="/appointment/delete/{{$record.Id}}" class="btn btn-danger">Delete</a>
          </td>
        </tr>
        {{end}}
      </tbody>
      <tfoot>
        <tr>
          <td colspan="4"><a class="btn btn-success" href="{{urlfor "AppointmentController.Add"}}">Add New Appointment</a></td>
        </tr>
      </tfoot>
    </table>
  </div>
</div>
