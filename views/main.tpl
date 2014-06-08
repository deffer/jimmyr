<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<link rel="stylesheet" type="text/css" href="/css/bootstrap-2.3.2.min.css"> 
	
</head>
<body>
<div class="container">

</div>
	<div name="leftColumn">
        <div class="well span4">
			<ul class="unstyled">
	            <li ng-repeat="b in selectedBookmarks">
	                <div>
	                    <i class="icon-remove pull-right" title="Delete" style="cursor: pointer;" ng-show="editMode" ng-click="deleteBookmark(b)"></i>
	                    <i class="icon-pencil pull-right" title="Edit" style="cursor: pointer;" ng-show="editMode" ng-click="editBookmark(b)"></i>
	                    <a href="{{b.url}}" target="_blank" style="font-size:small; display:block; margin-left: 2px; margin-right: 50px;">- {{b.title}}</a>
	                    <!--button type="button" class="close pull-right" aria-hidden="true" title="Delete">&times;</button-->
	                </div>
	            </li>
	        </ul>
		</div>
	</div>	
	
	<div name="rightColumn">
        <div class="well span6">
		<!--
			<%@include file="_utils.jsp" %>
		-->
		</div>
	</div>	
</body>
</html>