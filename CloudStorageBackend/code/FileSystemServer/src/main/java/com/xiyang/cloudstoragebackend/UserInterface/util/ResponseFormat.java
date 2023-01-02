package com.xiyang.cloudstoragebackend.UserInterface.util;

import net.minidev.json.JSONUtil;

import javax.servlet.http.HttpServletResponse;
import java.io.PrintWriter;

public class ResponseFormat {
    public static void writeJson(Object data, HttpServletResponse response){
        PrintWriter out=null;
        try{
            response.setCharacterEncoding("utf-8");
            response.setHeader("Content-Type","text/html;charset=utf-8");
            out=response.getWriter();
            //TODO you should complete the json encoder 
        }catch(Exception e){
            out.println("{result:false,message:'结果转换错误'}");
        }
    }
}
