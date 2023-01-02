package com.xiyang.cloudstoragebackend.ApplicationLayer;

import com.xiyang.cloudstoragebackend.DomainLayer.HDFSClient;
import org.apache.hadoop.fs.FileSystem;
import org.springframework.stereotype.Service;


import java.io.IOException;
import java.net.URISyntaxException;

@Service
public class FileOperation  {
   private FileSystem hdfs;
    FileOperation() throws URISyntaxException, IOException {
        hdfs= HDFSClient.getHDFSClient();
    }
    public  void uploadFileToHDFS(){

    }
    public void downloadFileFromHDFS(){

    }
    public void searchFolder(){

    }
}
