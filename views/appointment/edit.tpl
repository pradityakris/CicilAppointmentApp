</br>
</br>
</br>
</br>
<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Add Appointment</h1>
   </div>
 </div>
</div>
</div>
<div class="container">
  <div class="row">

    <p class="lead">Appointment Form</p>

    {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

    {{if .flash.notice}}
    <blockquote>{{.flash.notice}}</blockquote>
    {{end}}

    <p>
      <form role="form" id="user" method="POST">
        <div class="form-group {{if .Errors.Title}}has-error has-feedback{{end}}">
          <label for="title">Title： {{if .Errors.Title}}({{.Errors.Title}}){{end}}</label>
          <input name="title" type="text" value="{{.record.Title}}" class="form-control" tabindex="1" />
          {{if .Errors.Title}}<span class="glyphicon glyphicon-remove form-control-feedback"></span>{{end}}
        </div>

        <div class="form-group">
          <label for="date">Date：</label>
          <div class='input-group date' data-provide="datepicker">
          <input name="date" type="text" value="{{.record.Date}}" class="form-control" tabindex="2" />
          <span class="input-group-addon">
          <span class="glyphicon glyphicon-calendar"></span>
          </span>
          </div>
        </div>

        <div class="form-group">
          <label for="time">Time：</label>
          <div class="input-group bootstrap-timepicker timepicker">
          <input id="timepicker1" name="time" type="text" value="{{.record.Time}}" class="form-control" tabindex="3" />
          <span class="input-group-addon"><i class="glyphicon glyphicon-time"></i></span>
          </div>
        </div>

        <div class="form-group">
          <label for="desc">Description：</label>
          <input name="desc" type="textarea" value="{{.record.Description}}" class="form-control" tabindex="3" />
        </div>

        <div class="form-group">
          <label for="atten">Attendes：</label>
          <input name="atten" type="textarea" value="{{.record.Attendes}}" class="form-control" tabindex="4" />
        </div>

        <div class="form-group">
          <label for="loc">Location：</label>
          <input name="loc" type="text" value="{{.record.Location}}" class="form-control" tabindex="5" />
        </div>

        <input type="submit" value="Save Changes" class="btn btn-default" tabindex="6" /> &nbsp;
        <a href="#">Cancel</a>
      </form>
    </p>
  </div>
</div>
