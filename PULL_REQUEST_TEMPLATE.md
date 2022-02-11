## 📝  Description
#### 🗓 Estimated date: 31/03/2022
#### 📤 Partial deploy? <input type="radio" name="partial" id="yes" value="yes" checked="checked"/><label for="yes">Yes</label> <input type="radio" name="partial" id="no" value="no"/><label for="female">No</label>
- Please include a summary of the change or which issue is fixed. Also include relevant motivation and context. If it is a partial deploy, please include the mayor project and how this PR contributes to the whole feature/fix.


## 📑 Jira
- Please include the link to the Jira story related to these changes in the format:
  https://mercadolibre.atlassian.net/browse/FCNA-3623

## 📸  Screenshots
- Include screenshots if the changes impact directly our frontend

## 🚨  Dependencies
- List here any dependencies that are required for this change, for example:
- [x] mp-home-sections
- [ ] mp-home-shortcuts
- [ ] mp-home-proxy
- [ ] mp-home-layouts
- [ ] sections-navigation
- [ ] go-mpfcn-toolkit

## 👨🏽‍🔬  Testing
- Add here the used tests scopes and also a curl, so the reviewers can test and validate your changes. Provide instructions and please also list any relevant details for your test configuration

`Scope:` beta/gamma/delta

- Web
```
curl --location --request GET 'https://internal-api.mercadopago.com/mpmobile/home?caller.id=832385486&caller.siteId=MLM&client.id=6485336621976909'
```
- Mobile
```
curl --location --request GET 'https://internal-api.mercadopago.com/mpmobile/home?caller.id=832385486&caller.siteId=MLM&client.id=1311377052931992' \
--header 'User-Agent: MercadoPago-Android/2.206.0 (A910; Android 6.0; Build/MRA58K)'
```