# File Downloader from AWS S3 bucket and Digital Ocean Spaces

Since i have never found an optimal solution for downloading files form AWS S3 bucket and DigitalOcean Spaces i decide to create my own.
I hope that this solution would help some of you also. I tried to make it so that it can be easily reused by command line flags when running the program.
With this approach, you can easily create secrets in your GitHub repository and set them as the value of each of the flags in your pipeline.
By this approach you can securely upload your .env file in AWS S3 bucket or in DigitalOcean Spaces and in your pipeline you can download it.

The flags that this program supports are :
1. `-key` Access key for AWS/DigitalOcean bucket.
2. `-secret` Secret for AWS/DigitalOcean bucket.
3. `-endpoint` URL points to your s3/space. 
* Example `https://fra1.digitaloceanspaces.com/` or with the bucket name `https://examples.fra1.digitaloceanspaces.com/`.
4. `-bucket` Name of the bucket. 
* If you haven't added it with the url in the `-endpoint` flag you must provide it otherwise just skip this flag.
5. `-region` Region where your bucket is located
6. `-file` Name of the file that you want to download form the bucket.
7. `-destination` Absolute path and name of the file that you should save the output. 
* If you run the program in the folder that you desire to download the file you should just need to provide only the name otherwise you should provide the exact path.

With `go run downlad.go -h` you can see all flags and what you should add as value.

Examples for DigitalOcean Spaces :

1. This example is when you have added the bucket name in url of the `-endpoint` flag, also with the exact path where you want the file to be downloaded followed by the file name of the file (`src/.env`).

 * `go run download.go -key 123456 -secret secret1234 -endpoint https://examples.fra1.digitaloceanspaces.com/ -region fra1 -file .env -destination src/.env`

2. This example is when you used the `-bucket` flag to specify the bucket name and the url in the `-endpoint` flag doesn't contain the bucket name. Also in this example the following command would create the file in the directory where this program is run.
 * `go run download.go -key 123456 -secret secret1234 -bucket examples -endpoint https://fra1.digitaloceanspaces.com/ -region fra1 -file .env -destination .env`



The example for AWS S3 is analogical, because DigitalOcean Spaces uses their API.
If there is some mistakes you recognized in the code or in the readme, or you have questions you can write to me at `g.rudarska@gmail.com`