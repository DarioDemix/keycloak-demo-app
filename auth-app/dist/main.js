(()=>{const e=document.querySelector("#login-btn"),t=document.querySelector("#token-btn"),o="http://localhost:4000",n=`${o}/auth/realms/${process.env.REALM}/protocol/openid-connect`,c=`${n}/token`,r=`${n}/auth?response_type=code&client_id=auth-app&redirect_uri=${o}`,a={username:"myuser",password:"foobar",client_id:"auth-app",client_secret:"",grant_type:"password"};e.addEventListener("click",(()=>window.location.replace(r))),t.addEventListener("click",(()=>{const e=new URLSearchParams;for(const t in a)e.append(t,a[t]);fetch(c,{method:"POST",body:e}).then((e=>e.json())).then((e=>console.log(e)))}))})();