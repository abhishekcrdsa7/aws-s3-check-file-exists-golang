package main

import (
	"fmt"
	"time"
	"github.com/AdRoll/goamz/aws"
	"github.com/AdRoll/goamz/s3"
)

func main() {
	now := time.Now()
	threeDays := time.Hour * 24 * 3
	diff := now.Add(threeDays) //expiration date of STS token
	//if you dont have STS token don't worry it will still work
	//just put any date in that case of type time.Time

	//leave the token as "" if you don't have and you are not concerned with AWS STS
	auth, _ := aws.GetAuth("<AWS ACCESS KEY>", "<AWS SECRET ACCESS KEY", "<TOKEN>", <expiration date time.Time>) 
	//return auth, error
	
	region := aws.GetRegion("<Region in which you created s3 bucket>") //example us-east-1 or ap-south-1
	//change the signature to new v4signature in the s3.go file in the package
	
	S3 := s3.New(auth, region) //creates new S3
	bucket := S3.Bucket("<Bucket Name>") //example myBucket 
	exists, _ := bucket.Exists("<Path to file in bucket>") //example /index.html or /mypicture.jpg 
	//returns true/false, error
	fmt.Println(exists)
}