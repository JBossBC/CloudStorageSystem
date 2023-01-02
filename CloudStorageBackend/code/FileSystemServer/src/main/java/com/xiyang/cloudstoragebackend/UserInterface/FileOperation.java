package com.xiyang.cloudstoragebackend.UserInterface;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

@Controller
@RequestMapping("/file/")
public class FileOperation {
   @RequestMapping(method = RequestMethod.POST,value = "/upload")
    public  void uploadFileToHDFS(@RequestParam(value="file",required = true) MultipartFile file, HttpServletRequest request, HttpServletResponse response){
    }
    @RequestMapping(method = RequestMethod.GET,value = "/download")
    public void downloadFileFromHDFS(HttpServletRequest request, HttpServletResponse response){



    }
    @RequestMapping(method = RequestMethod.GET,value = "/search")
    public void searchFilesByHDFS(HttpServletRequest request, HttpServletResponse response){

    }

}
