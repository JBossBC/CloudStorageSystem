package com.xiyang.cloudstoragebackend.UserInterface.util;

import java.util.HashMap;
import java.util.Map;

public class ResultCodeInfo {
    public static final String RESULT_SUCCEED="true";
    public static final String RESULT_FAIL="false";
    public static final String WORD_RESULT="result";
    public static final String WORD_MESSAGE="message";
    public ResultCodeInfo(){

    }
    public static Map<String,Object> succeed(String message){
        Map<String,Object> result=new HashMap<>();
        result.put("result","true");
        if(message==null){
            message="操作成功";
        }else {
            result.put("message", message);
        }
        return result;
    }
    public static Map<String, Object> fail(String message) {


        Map<String, Object> ret = new HashMap();
        ret.put("result", "false");
        if (message == null) {
            ret.put("message","操作失败");
        }else {
            ret.put("message", message);
        }
        return ret;
    }

}
