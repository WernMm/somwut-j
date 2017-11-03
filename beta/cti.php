<?php
?>
<html>
	<head>
	<title>Test Open CTI</title>
   	<!-- <script type="text/javascript" src="http://domain:port/support/api/41.0/lightning/opencti_min.js"></script> -->
	<script type="text/javascript" src="https://somwut-j-opencti-dev-ed--c.ap5.visual.force.com/support/api/41.0/lightning/opencti_min.js" type="text/javascript"></script>
   <script type="text/javascript">							
 	var callback = function(response) {
         if (response.success) {
            console.log('API method call executed successfully! returnValue:', response.returnValue);
         } else { 
            console.error('Something went wrong! Errors:', response.errors);
         }
      	};
   	function screenPop(scopeName) {
    	sforce.opencti.screenPop({type: sforce.opencti.SCREENPOP_TYPE.OBJECTHOME, params: {scope:scopeName}, callback: callback  });
	}

	function screenPopRecord(recordIdValue) {
		sforce.opencti.screenPop({type: sforce.opencti.SCREENPOP_TYPE.SOBJECT, params: {
			recordId: recordIdValue
		}});		
	}
</script>
</head>
<body>
       <button onclick="screenPop('Contact');">pop contacts</button>
       <button onclick="screenPop('Case');">pop cases</button>
       <button onclick="screenPopRecord('0037F00000CTnosQAD');">pop record</button>

</body>
</html>

