(window.webpackJsonp=window.webpackJsonp||[]).push([[13],{120:function(e,t,n){"use strict";n.d(t,"a",(function(){return i}));var r=n(2),i=e=>({type:r.A,boolean:e})},128:function(e,t,n){"use strict";t.a={en:{Login:{loginToAccount:"Login to your account",username:"Username",password:"Password",login:"login",back:"back"},Menu:{home:"Home",channels:"TV channels",movies:"Movies",series:"Series",radio:"Radio",settings:"Settings"},Home:{channels:"TV Channels",newMovies:"New movies",newSeries:"New series"},Settings:{userInfo:"User info",seeYourInfo:"See your info",speedTest:"Speed test",checkSpeed:"Check your internet speed",pinCode:"Pin code",changePin:"Change your PIN code",lockedCategories:"Locked categories",chooseLockedCategories:"Choose locked categories",language:"Language",chooseLanguage:"Choose app language",logOut:"Log out",logOutFromProfile:"Log out from this profile"}},hr:{Login:{loginToAccount:"Ulogujte se na akaunt",username:"Korisnicko ime",password:"Sifra",login:"uloguj se",back:"nazad"},Menu:{home:"Pocetna",channels:"TV kanali",movies:"Filmovi",series:"Serije",radio:"Radio",settings:"Podesavanja"},Home:{channels:"TV Kanali",newMovies:"Novi filmovi",newSeries:"Nove serije"},Settings:{userInfo:"Info korisnika",seeYourInfo:"Pogledajte vase podatke",speedTest:"Brzina interneta",checkSpeed:"Proverite vasu brzinu interneta",pinCode:"PIN kod",changePin:"Promenite vas PIN kod",lockedCategories:"Zakljucane kategorije",chooseLockedCategories:"Izaberite zakljucane kategorije",language:"Jezike",chooseLanguage:"Izaberite jezik aplikacije",logOut:"Izloguj se",logOutFromProfile:"Odjavite se"}},de:{Login:{loginToAccount:"Melde dich in deinem Konto an",username:"Nutzername",password:"Passwort",login:"anmeldung",back:"zurück"},Menu:{home:"Zuhause",channels:"Fernsehsender",movies:"Filme",series:"Serie",radio:"Radio",settings:"Einstellungen"},Home:{channels:"Fernsehsender",newMovies:"Neue Filme",newSeries:"Neue Serie"},Settings:{userInfo:"Benutzerinformation",seeYourInfo:"Sehen Sie Ihre Informationen",speedTest:"Geschwindigkeitstest",checkSpeed:"Überprüfen Sie Ihre Internetgeschwindigkeit",pinCode:"Geheimzahl",changePin:"Ändern Sie Ihren PIN-Code",lockedCategories:"Gesperrte Kategorien",chooseLockedCategories:"Wählen Sie gesperrte Kategorien",language:"Sprache",chooseLanguage:"Wählen Sie die App-Sprache",logOut:"Ausloggen",logOutFromProfile:"Melden Sie sich vom Profil ab"}}}},142:function(e,t,n){e.exports=function(){"use strict";var e="millisecond",t="second",n="minute",r="hour",i="day",a="week",o="month",s="quarter",u="year",c="date",l=/^(\d{4})[-/]?(\d{1,2})?[-/]?(\d{0,2})[^0-9]*(\d{1,2})?:?(\d{1,2})?:?(\d{1,2})?[.:]?(\d+)?$/,h=/\[([^\]]+)]|Y{1,4}|M{1,4}|D{1,2}|d{1,4}|H{1,2}|h{1,2}|a|A|m{1,2}|s{1,2}|Z{1,2}|SSS/g,f={name:"en",weekdays:"Sunday_Monday_Tuesday_Wednesday_Thursday_Friday_Saturday".split("_"),months:"January_February_March_April_May_June_July_August_September_October_November_December".split("_")},d=function(e,t,n){var r=String(e);return!r||r.length>=t?e:""+Array(t+1-r.length).join(n)+e},m={s:d,z:function(e){var t=-e.utcOffset(),n=Math.abs(t),r=Math.floor(n/60),i=n%60;return(t<=0?"+":"-")+d(r,2,"0")+":"+d(i,2,"0")},m:function e(t,n){if(t.date()<n.date())return-e(n,t);var r=12*(n.year()-t.year())+(n.month()-t.month()),i=t.clone().add(r,o),a=n-i<0,s=t.clone().add(r+(a?-1:1),o);return+(-(r+(n-i)/(a?i-s:s-i))||0)},a:function(e){return e<0?Math.ceil(e)||0:Math.floor(e)},p:function(l){return{M:o,y:u,w:a,d:i,D:c,h:r,m:n,s:t,ms:e,Q:s}[l]||String(l||"").toLowerCase().replace(/s$/,"")},u:function(e){return void 0===e}},g="en",p={};p[g]=f;var v=function(e){return e instanceof w},y=function(e,t,n){var r;if(!e)return g;if("string"==typeof e)p[e]&&(r=e),t&&(p[e]=t,r=e);else{var i=e.name;p[i]=e,r=i}return!n&&r&&(g=r),r||!n&&g},b=function(e,t){if(v(e))return e.clone();var n="object"==typeof t?t:{};return n.date=e,n.args=arguments,new w(n)},M=m;M.l=y,M.i=v,M.w=function(e,t){return b(e,{locale:t.$L,utc:t.$u,x:t.$x,$offset:t.$offset})};var w=function(){function f(e){this.$L=y(e.locale,null,!0),this.parse(e)}var d=f.prototype;return d.parse=function(e){this.$d=function(e){var t=e.date,n=e.utc;if(null===t)return new Date(NaN);if(M.u(t))return new Date;if(t instanceof Date)return new Date(t);if("string"==typeof t&&!/Z$/i.test(t)){var r=t.match(l);if(r){var i=r[2]-1||0,a=(r[7]||"0").substring(0,3);return n?new Date(Date.UTC(r[1],i,r[3]||1,r[4]||0,r[5]||0,r[6]||0,a)):new Date(r[1],i,r[3]||1,r[4]||0,r[5]||0,r[6]||0,a)}}return new Date(t)}(e),this.$x=e.x||{},this.init()},d.init=function(){var e=this.$d;this.$y=e.getFullYear(),this.$M=e.getMonth(),this.$D=e.getDate(),this.$W=e.getDay(),this.$H=e.getHours(),this.$m=e.getMinutes(),this.$s=e.getSeconds(),this.$ms=e.getMilliseconds()},d.$utils=function(){return M},d.isValid=function(){return!("Invalid Date"===this.$d.toString())},d.isSame=function(e,t){var n=b(e);return this.startOf(t)<=n&&n<=this.endOf(t)},d.isAfter=function(e,t){return b(e)<this.startOf(t)},d.isBefore=function(e,t){return this.endOf(t)<b(e)},d.$g=function(e,t,n){return M.u(e)?this[t]:this.set(n,e)},d.unix=function(){return Math.floor(this.valueOf()/1e3)},d.valueOf=function(){return this.$d.getTime()},d.startOf=function(e,s){var l=this,h=!!M.u(s)||s,f=M.p(e),d=function(e,t){var n=M.w(l.$u?Date.UTC(l.$y,t,e):new Date(l.$y,t,e),l);return h?n:n.endOf(i)},m=function(e,t){return M.w(l.toDate()[e].apply(l.toDate("s"),(h?[0,0,0,0]:[23,59,59,999]).slice(t)),l)},g=this.$W,p=this.$M,v=this.$D,y="set"+(this.$u?"UTC":"");switch(f){case u:return h?d(1,0):d(31,11);case o:return h?d(1,p):d(0,p+1);case a:var b=this.$locale().weekStart||0,w=(g<b?g+7:g)-b;return d(h?v-w:v+(6-w),p);case i:case c:return m(y+"Hours",0);case r:return m(y+"Minutes",1);case n:return m(y+"Seconds",2);case t:return m(y+"Milliseconds",3);default:return this.clone()}},d.endOf=function(e){return this.startOf(e,!1)},d.$set=function(a,s){var l,h=M.p(a),f="set"+(this.$u?"UTC":""),d=(l={},l[i]=f+"Date",l[c]=f+"Date",l[o]=f+"Month",l[u]=f+"FullYear",l[r]=f+"Hours",l[n]=f+"Minutes",l[t]=f+"Seconds",l[e]=f+"Milliseconds",l)[h],m=h===i?this.$D+(s-this.$W):s;if(h===o||h===u){var g=this.clone().set(c,1);g.$d[d](m),g.init(),this.$d=g.set(c,Math.min(this.$D,g.daysInMonth())).$d}else d&&this.$d[d](m);return this.init(),this},d.set=function(e,t){return this.clone().$set(e,t)},d.get=function(e){return this[M.p(e)]()},d.add=function(e,s){var c,l=this;e=Number(e);var h=M.p(s),f=function(t){var n=b(l);return M.w(n.date(n.date()+Math.round(t*e)),l)};if(h===o)return this.set(o,this.$M+e);if(h===u)return this.set(u,this.$y+e);if(h===i)return f(1);if(h===a)return f(7);var d=(c={},c[n]=6e4,c[r]=36e5,c[t]=1e3,c)[h]||1,m=this.$d.getTime()+e*d;return M.w(m,this)},d.subtract=function(e,t){return this.add(-1*e,t)},d.format=function(e){var t=this;if(!this.isValid())return"Invalid Date";var n=e||"YYYY-MM-DDTHH:mm:ssZ",r=M.z(this),i=this.$locale(),a=this.$H,o=this.$m,s=this.$M,u=i.weekdays,c=i.months,l=function(e,r,i,a){return e&&(e[r]||e(t,n))||i[r].substr(0,a)},f=function(e){return M.s(a%12||12,e,"0")},d=i.meridiem||function(e,t,n){var r=e<12?"AM":"PM";return n?r.toLowerCase():r},m={YY:String(this.$y).slice(-2),YYYY:this.$y,M:s+1,MM:M.s(s+1,2,"0"),MMM:l(i.monthsShort,s,c,3),MMMM:l(c,s),D:this.$D,DD:M.s(this.$D,2,"0"),d:String(this.$W),dd:l(i.weekdaysMin,this.$W,u,2),ddd:l(i.weekdaysShort,this.$W,u,3),dddd:u[this.$W],H:String(a),HH:M.s(a,2,"0"),h:f(1),hh:f(2),a:d(a,o,!0),A:d(a,o,!1),m:String(o),mm:M.s(o,2,"0"),s:String(this.$s),ss:M.s(this.$s,2,"0"),SSS:M.s(this.$ms,3,"0"),Z:r};return n.replace(h,(function(e,t){return t||m[e]||r.replace(":","")}))},d.utcOffset=function(){return 15*-Math.round(this.$d.getTimezoneOffset()/15)},d.diff=function(e,c,l){var h,f=M.p(c),d=b(e),m=6e4*(d.utcOffset()-this.utcOffset()),g=this-d,p=M.m(this,d);return p=(h={},h[u]=p/12,h[o]=p,h[s]=p/3,h[a]=(g-m)/6048e5,h[i]=(g-m)/864e5,h[r]=g/36e5,h[n]=g/6e4,h[t]=g/1e3,h)[f]||g,l?p:M.a(p)},d.daysInMonth=function(){return this.endOf(o).$D},d.$locale=function(){return p[this.$L]},d.locale=function(e,t){if(!e)return this.$L;var n=this.clone(),r=y(e,t,!0);return r&&(n.$L=r),n},d.clone=function(){return M.w(this.$d,this)},d.toDate=function(){return new Date(this.valueOf())},d.toJSON=function(){return this.isValid()?this.toISOString():null},d.toISOString=function(){return this.$d.toISOString()},d.toString=function(){return this.$d.toUTCString()},f}(),S=w.prototype;return b.prototype=S,[["$ms",e],["$s",t],["$m",n],["$H",r],["$W",i],["$M",o],["$y",u],["$D",c]].forEach((function(e){S[e[1]]=function(t){return this.$g(t,e[0],e[1])}})),b.extend=function(e,t){return e.$i||(e(t,w,b),e.$i=!0),b},b.locale=y,b.isDayjs=v,b.unix=function(e){return b(1e3*e)},b.en=p[g],b.Ls=p,b.p={},b}()},147:function(e,t,n){(t=n(41)(!1)).push([e.i,".Main{display:-webkit-flex;height:100%}.Main .view{width:100%;overflow:hidden;background:radial-gradient(29.26% 29.26% at 50% 56.06%, #481251 0%, #1D1D1D 100%);z-index:100}.Main .view .header{position:fixed;top:0;width:100%;z-index:10000;height:137px;overflow:hidden}\n",""]),e.exports=t},148:function(e,t,n){(t=n(41)(!1)).push([e.i,".Menu{position:absolute;background:linear-gradient(#1E0F29, #251E2B);width:370px;height:100%;overflow:hidden;z-index:102;transition:.2s;padding-top:180px;box-sizing:border-box}.Menu .logo{position:absolute;top:0;left:50%;transform:translateX(-50%);transition:.2s}.Menu .nav-item{display:-webkit-flex;align-items:center;width:370px;height:120px;font-size:25px;font-weight:600;margin-bottom:20px;text-transform:uppercase}.Menu .nav-item img{margin-left:23px;margin-right:26px}.Menu .nav-item.active{background:#573076}.Menu .time{position:absolute;bottom:10px;font-size:25px;left:20px;width:370px}.Menu .time span{margin-right:95px}.Menu.close{width:95px}.Menu.close .logo{transform:translateX(-190px)}.Menu.close .nav-item p{opacity:0}@media (max-width: 1280px){.Menu{width:250px}.Menu .logo img{width:200px}.Menu .nav-item{height:50px;margin-bottom:40px;font-size:17px}.Menu .nav-item img{width:38px;margin-left:12px;margin-right:15px}.Menu.close{width:70px}}\n",""]),e.exports=t},194:function(e,t,n){var r=n(147);"string"==typeof r&&(r=[[e.i,r,""]]);var i={hmr:!0,transform:void 0,insertInto:void 0},a=n(42)(r,i);r.locals&&(e.exports=r.locals),e.hot.accept(147,(function(){var t=n(147);if("string"==typeof t&&(t=[[e.i,t,""]]),!function(e,t){var n,r=0;for(n in e){if(!t||e[n]!==t[n])return!1;r++}for(n in t)r--;return 0===r}(r.locals,t.locals))throw new Error("Aborting CSS HMR due to changed css-modules locals.");a(t)})),e.hot.dispose((function(){a()}))},195:function(e,t,n){var r=n(148);"string"==typeof r&&(r=[[e.i,r,""]]);var i={hmr:!0,transform:void 0,insertInto:void 0},a=n(42)(r,i);r.locals&&(e.exports=r.locals),e.hot.accept(148,(function(){var t=n(148);if("string"==typeof t&&(t=[[e.i,t,""]]),!function(e,t){var n,r=0;for(n in e){if(!t||e[n]!==t[n])return!1;r++}for(n in t)r--;return 0===r}(r.locals,t.locals))throw new Error("Aborting CSS HMR due to changed css-modules locals.");a(t)})),e.hot.dispose((function(){a()}))},236:function(e,t,n){"use strict";n.r(t);var r=n(1),i=n.n(r),a=(n(194),n(195),n(8)),o=n(24),s=n(54),u=n(120),c=n(142),l=n.n(c);function h(e,t){return function(e){if(Array.isArray(e))return e}(e)||function(e,t){if("undefined"==typeof Symbol||!(Symbol.iterator in Object(e)))return;var n=[],r=!0,i=!1,a=void 0;try{for(var o,s=e[Symbol.iterator]();!(r=(o=s.next()).done)&&(n.push(o.value),!t||n.length!==t);r=!0);}catch(e){i=!0,a=e}finally{try{r||null==s.return||s.return()}finally{if(i)throw a}}return n}(e,t)||function(e,t){if(!e)return;if("string"==typeof e)return f(e,t);var n=Object.prototype.toString.call(e).slice(8,-1);"Object"===n&&e.constructor&&(n=e.constructor.name);if("Map"===n||"Set"===n)return Array.from(e);if("Arguments"===n||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n))return f(e,t)}(e,t)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function f(e,t){(null==t||t>e.length)&&(t=e.length);for(var n=0,r=new Array(t);n<t;n++)r[n]=e[n];return r}function d(){var e=h(Object(r.useState)((new Date).getTime()),2),t=e[0],n=e[1];return Object(r.useEffect)(()=>{var e=setInterval(()=>{n((new Date).getTime())},6e4);return()=>clearInterval(e)},[t]),i.a.createElement(i.a.Fragment,null,i.a.createElement("span",null,l()(t).format("HH:mm")),l()(t).format("DD MMMM YYYY"))}var m=n(128);function g(e,t){return function(e){if(Array.isArray(e))return e}(e)||function(e,t){if("undefined"==typeof Symbol||!(Symbol.iterator in Object(e)))return;var n=[],r=!0,i=!1,a=void 0;try{for(var o,s=e[Symbol.iterator]();!(r=(o=s.next()).done)&&(n.push(o.value),!t||n.length!==t);r=!0);}catch(e){i=!0,a=e}finally{try{r||null==s.return||s.return()}finally{if(i)throw a}}return n}(e,t)||function(e,t){if(!e)return;if("string"==typeof e)return p(e,t);var n=Object.prototype.toString.call(e).slice(8,-1);"Object"===n&&e.constructor&&(n=e.constructor.name);if("Map"===n||"Set"===n)return Array.from(e);if("Arguments"===n||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n))return p(e,t)}(e,t)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function p(e,t){(null==t||t>e.length)&&(t=e.length);for(var n=0,r=new Array(t);n<t;n++)r[n]=e[n];return r}var v=e=>{var t=g(Object(r.useState)(0),2),n=t[0],a=t[1];return Object(r.useEffect)(()=>{var t=()=>{switch(n){case 0:e.toggleMenu(!1),e.history.push("/home");break;case 1:e.toggleMenu(!1),e.history.push("/tv-kanali");break;case 2:e.toggleMenu(!1),e.history.push("/filmovi");break;case 3:e.toggleMenu(!1),e.history.push("/serije");break;case 4:e.toggleMenu(!1),e.history.push("/radio");break;case 5:e.toggleMenu(!1),e.history.push("/podesavanja")}},r=e=>{switch(e.keyCode){case s.a.down:n<5&&a(e=>e+1);break;case s.a.up:n>0&&a(e=>e-1);break;case s.a.enter:case s.a.right:t()}};return e.menu?document.addEventListener("keydown",r):document.removeEventListener("keydown",r),()=>document.removeEventListener("keydown",r)},[n,e]),i.a.createElement("div",{className:"Menu"+(e.menu?"":" close")},i.a.createElement("div",{className:"logo"},i.a.createElement("img",{src:"http://fixanything.boopro.mycpanel.rs/Layout/logo.png",alt:"MaxiTV"})),i.a.createElement("div",{className:"nav-item"+(0===n?" active":"")},i.a.createElement("img",{src:"http://fixanything.boopro.mycpanel.rs/Layout/pocetna.png",alt:"Pocetna"}),i.a.createElement("p",null,m.a[e.ln].Menu.home)),i.a.createElement("div",{className:"nav-item"+(1===n?" active":"")},i.a.createElement("img",{src:"http://fixanything.boopro.mycpanel.rs/Layout/tvkanali.png",alt:"Pocetna"}),i.a.createElement("p",null,m.a[e.ln].Menu.channels)),i.a.createElement("div",{className:"nav-item"+(2===n?" active":"")},i.a.createElement("img",{src:"http://fixanything.boopro.mycpanel.rs/Layout/filmovi.png",alt:"Filmovi"}),i.a.createElement("p",null,m.a[e.ln].Menu.movies)),i.a.createElement("div",{className:"nav-item"+(3===n?" active":"")},i.a.createElement("img",{src:"http://fixanything.boopro.mycpanel.rs/Layout/serije.png",alt:"Pocetna"}),i.a.createElement("p",null,m.a[e.ln].Menu.series)),i.a.createElement("div",{className:"nav-item"+(4===n?" active":"")},i.a.createElement("img",{src:"http://fixanything.boopro.mycpanel.rs/Layout/radio.png",alt:"Pocetna"}),i.a.createElement("p",null,m.a[e.ln].Menu.radio)),i.a.createElement("div",{className:"nav-item"+(5===n?" active":"")},i.a.createElement("img",{src:"http://fixanything.boopro.mycpanel.rs/Layout/podesavanja.png",alt:"Pocetna"}),i.a.createElement("p",null,m.a[e.ln].Menu.settings)),i.a.createElement("div",{className:"time"},i.a.createElement(d,null)))},y=Object(o.b)(e=>({menu:e.menu.menu,ln:e.language}),{toggleMenu:u.a})(Object(a.g)(i.a.memo(v))),b=n(33);function M(e,t){return function(e){if(Array.isArray(e))return e}(e)||function(e,t){if("undefined"==typeof Symbol||!(Symbol.iterator in Object(e)))return;var n=[],r=!0,i=!1,a=void 0;try{for(var o,s=e[Symbol.iterator]();!(r=(o=s.next()).done)&&(n.push(o.value),!t||n.length!==t);r=!0);}catch(e){i=!0,a=e}finally{try{r||null==s.return||s.return()}finally{if(i)throw a}}return n}(e,t)||function(e,t){if(!e)return;if("string"==typeof e)return w(e,t);var n=Object.prototype.toString.call(e).slice(8,-1);"Object"===n&&e.constructor&&(n=e.constructor.name);if("Map"===n||"Set"===n)return Array.from(e);if("Arguments"===n||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n))return w(e,t)}(e,t)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function w(e,t){(null==t||t>e.length)&&(t=e.length);for(var n=0,r=new Array(t);n<t;n++)r[n]=e[n];return r}var S=e=>{var t=M(Object(r.useState)(!0),2),n=t[0],a=t[1],o=M(Object(r.useState)([]),2),s=o[0],u=o[1],c=M(Object(r.useState)(0),2),l=c[0],h=c[1];return Object(r.useEffect)(()=>{var e=setInterval(()=>{s.length-1>l?h(e=>e+1):h(0)},15e3);return()=>clearInterval(e)},[s,l]),Object(r.useEffect)(()=>{b.a.get("/banners").then(({data:e})=>{var t=e.banners.map(e=>e.image_url);u(t)})},[]),Object(r.useEffect)(()=>{var t=e.location.pathname;a("/home"===t||"/filmovi"===t||"/serije"===t||"/podesavanja"===t)},[e.location]),i.a.createElement("div",{className:"Main"},i.a.createElement(y,{open:e.menu}),i.a.createElement("div",{className:"view"},n&&i.a.createElement("div",{className:"header",style:{background:"url(".concat(s[l],")")}}),e.children))};t.default=Object(a.g)(i.a.memo(S))}}]);