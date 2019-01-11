package userinfo

/*
#include <stdio.h>
#include <unistd.h>     // getuid
#include <pwd.h>        // getpwuid
#include <sys/types.h>  // uid_t

char* getuserinfo() {
   uid_t          user_id;
   struct passwd *user_pw;

   user_id  = getuid();             // 사용자 아이디를 구하고
   user_pw  = getpwuid( user_id);   // 아이디로 사용자 정보 구하기

   // printf( "user name      :%s\n", user_pw->pw_name  );
   // printf( "real name      :%s\n", user_pw->pw_gecos );
   return user_pw->pw_name;
}
*/
import "C"

type Testlib struct{}

func (testlib *Testlib) GetInfo(v interface{}, result *string) error {
	*result = C.GoString(C.getuserinfo())
	return nil
}
