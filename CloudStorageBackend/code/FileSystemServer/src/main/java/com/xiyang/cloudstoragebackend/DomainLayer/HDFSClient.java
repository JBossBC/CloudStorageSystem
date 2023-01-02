package com.xiyang.cloudstoragebackend.DomainLayer;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.FileSystem;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;

public class HDFSClient {
    private static FileSystem hdfs;
    private static Configuration config;

    public synchronized static FileSystem getHDFSClient() throws URISyntaxException, IOException {
             if(hdfs==null){
                 synchronized (HDFSClient.class){
                      if(hdfs==null){
                          hdfs=FileSystem.newInstance(new URI("hdfs://master:9000"),config);
                      }
                 }
             }
             return hdfs;
    }

}
