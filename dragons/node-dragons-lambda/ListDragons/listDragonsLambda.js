// # Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// #
// # Licensed under the Apache License, Version 2.0 (the "License"). You may not use this file except
// # in compliance with the License. A copy of the License is located at
// #
// # https://aws.amazon.com/apache-2-0/
// #
// # or in the "license" file accompanying this file. This file is distributed on an "AS IS" BASIS,
// # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// # specific language governing permissions and limitations under the License.

var AWSXRay = require('aws-xray-sdk')
var AWS = AWSXRay.captureAWS(require('aws-sdk'));

const s3 = new AWS.S3({
    region: 'us-east-1'
});

const ssm = new AWS.SSM({
    region: 'us-east-1'
});


exports.handler =  function(event, context, callback) {
    readDragons(event,callback);
}

async function readDragons(event, callback) {
  
   var fileName = await getFileName();
   var bucketName = await getBucketName();
   var dragonData = readDragonsFromS3(bucketName, fileName, event, callback);
   
}

async function getFileName() {
 
    var fileNameParams = {
        Name: 'dragon_data_file_name', 
        WithDecryption: false
    };
    var promise = await ssm.getParameter(fileNameParams).promise();
    return promise.Parameter.Value;
}

async function getBucketName() {
  var bucketNameParams = {
    Name: 'dragon_data_bucket_name', 
    WithDecryption: false
   };
  
  var promise = await ssm.getParameter(bucketNameParams).promise();
  return promise.Parameter.Value;

}

function readDragonsFromS3(bucketName, fileName, event, callback) {
    
    var expression = "select * from S3Object[*][*] s"

    if(event.queryStringParameters && event.queryStringParameters.dragonName && event.queryStringParameters.dragonName !== "")
            expression = "select * from S3Object[*][*] s where s.dragon_name_str =  '" + event["queryStringParameters"]['dragonName'] + "'"
    else if (event.queryStringParameters && event.queryStringParameters.family && event.queryStringParameters.family !== "")
            expression = "select * from S3Object[*][*] s where s.family_str =  '" + event["queryStringParameters"]['family'] + "'"
            
  s3.selectObjectContent({
     Bucket: bucketName,
     Expression: expression,
     ExpressionType: 'SQL',
     Key: fileName,
     InputSerialization: {
       JSON: {
         Type: 'DOCUMENT',
       }
     },
     OutputSerialization: {
       JSON: {
         RecordDelimiter: ','
       }
     }
   }, function(err, data) {
     if (err) {
         console.log(err);
     } else {
       return handleData(data, callback);
     }
   }
  );
}

function handleData(data, callback) {
    var event = data.Payload;
    const resultStream = [];

    event.on('data', function (event) {
        if (event.Records) {
            if (event.Records) {
                resultStream.push(event.Records.Payload);
            }
        }
    });

    event.on("end", function () {
        let recordsString = Buffer.concat(resultStream).toString('utf8');
        // remove any trailing commas
        recordsString = recordsString.replace(/\,$/, '');
        // parse as an array
        recordsString = `[${recordsString}]`;
        let records = JSON.parse(recordsString);

        callback(null, {
            "statusCode": 200,
            "body": JSON.stringify(records),
            "headers": { "access-control-allow-origin": "*" }
        })
    });
}
