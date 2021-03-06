package vpn

type Vpnsessionaction struct {
	Advancedclientlessvpnmode  string      `json:"advancedclientlessvpnmode,omitempty"`
	Allowedlogingroups         string      `json:"allowedlogingroups,omitempty"`
	Allprotocolproxy           string      `json:"allprotocolproxy,omitempty"`
	Alwaysonprofilename        string      `json:"alwaysonprofilename,omitempty"`
	Authorizationgroup         string      `json:"authorizationgroup,omitempty"`
	Autoproxyurl               string      `json:"autoproxyurl,omitempty"`
	Builtin                    interface{} `json:"builtin,omitempty"`
	Citrixreceiverhome         string      `json:"citrixreceiverhome,omitempty"`
	Clientchoices              string      `json:"clientchoices,omitempty"`
	Clientcleanupprompt        string      `json:"clientcleanupprompt,omitempty"`
	Clientconfiguration        interface{} `json:"clientconfiguration,omitempty"`
	Clientdebug                string      `json:"clientdebug,omitempty"`
	Clientidletimeout          int         `json:"clientidletimeout,omitempty"`
	Clientidletimeoutwarning   int         `json:"clientidletimeoutwarning,omitempty"`
	Clientlessmodeurlencoding  string      `json:"clientlessmodeurlencoding,omitempty"`
	Clientlesspersistentcookie string      `json:"clientlesspersistentcookie,omitempty"`
	Clientlessvpnmode          string      `json:"clientlessvpnmode,omitempty"`
	Clientoptions              string      `json:"clientoptions,omitempty"`
	Clientsecurity             string      `json:"clientsecurity,omitempty"`
	Clientsecuritygroup        string      `json:"clientsecuritygroup,omitempty"`
	Clientsecuritylog          string      `json:"clientsecuritylog,omitempty"`
	Clientsecuritymessage      string      `json:"clientsecuritymessage,omitempty"`
	Defaultauthorizationaction string      `json:"defaultauthorizationaction,omitempty"`
	Dnsvservername             string      `json:"dnsvservername,omitempty"`
	Emailhome                  string      `json:"emailhome,omitempty"`
	Epaclienttype              string      `json:"epaclienttype,omitempty"`
	Forcecleanup               interface{} `json:"forcecleanup,omitempty"`
	Forcedtimeout              int         `json:"forcedtimeout,omitempty"`
	Forcedtimeoutwarning       int         `json:"forcedtimeoutwarning,omitempty"`
	Ftpproxy                   string      `json:"ftpproxy,omitempty"`
	Gopherproxy                string      `json:"gopherproxy,omitempty"`
	Homepage                   string      `json:"homepage,omitempty"`
	Httpport                   interface{} `json:"httpport,omitempty"`
	Httpproxy                  string      `json:"httpproxy,omitempty"`
	Icaproxy                   string      `json:"icaproxy,omitempty"`
	Iconwithreceiver           string      `json:"iconwithreceiver,omitempty"`
	Iipdnssuffix               string      `json:"iipdnssuffix,omitempty"`
	Kcdaccount                 string      `json:"kcdaccount,omitempty"`
	Killconnections            string      `json:"killconnections,omitempty"`
	Linuxpluginupgrade         string      `json:"linuxpluginupgrade,omitempty"`
	Locallanaccess             string      `json:"locallanaccess,omitempty"`
	Loginscript                string      `json:"loginscript,omitempty"`
	Logoutscript               string      `json:"logoutscript,omitempty"`
	Macpluginupgrade           string      `json:"macpluginupgrade,omitempty"`
	Name                       string      `json:"name,omitempty"`
	Ntdomain                   string      `json:"ntdomain,omitempty"`
	Pcoipprofilename           string      `json:"pcoipprofilename,omitempty"`
	Proxy                      string      `json:"proxy,omitempty"`
	Proxyexception             string      `json:"proxyexception,omitempty"`
	Proxylocalbypass           string      `json:"proxylocalbypass,omitempty"`
	Rdpclientprofilename       string      `json:"rdpclientprofilename,omitempty"`
	Rfc1918                    string      `json:"rfc1918,omitempty"`
	Securebrowse               string      `json:"securebrowse,omitempty"`
	Sesstimeout                int         `json:"sesstimeout,omitempty"`
	Sfgatewayauthtype          string      `json:"sfgatewayauthtype,omitempty"`
	Smartgroup                 string      `json:"smartgroup,omitempty"`
	Socksproxy                 string      `json:"socksproxy,omitempty"`
	Splitdns                   string      `json:"splitdns,omitempty"`
	Splittunnel                string      `json:"splittunnel,omitempty"`
	Spoofiip                   string      `json:"spoofiip,omitempty"`
	Sslproxy                   string      `json:"sslproxy,omitempty"`
	Sso                        string      `json:"sso,omitempty"`
	Ssocredential              string      `json:"ssocredential,omitempty"`
	Storefronturl              string      `json:"storefronturl,omitempty"`
	Transparentinterception    string      `json:"transparentinterception,omitempty"`
	Useiip                     string      `json:"useiip,omitempty"`
	Usemip                     string      `json:"usemip,omitempty"`
	Useraccounting             string      `json:"useraccounting,omitempty"`
	Wihome                     string      `json:"wihome,omitempty"`
	Wihomeaddresstype          string      `json:"wihomeaddresstype,omitempty"`
	Windowsautologon           string      `json:"windowsautologon,omitempty"`
	Windowsclienttype          string      `json:"windowsclienttype,omitempty"`
	Windowspluginupgrade       string      `json:"windowspluginupgrade,omitempty"`
	Winsip                     string      `json:"winsip,omitempty"`
	Wiportalmode               string      `json:"wiportalmode,omitempty"`
}
