import i18n from 'i18next';
// Language Files
import en from './en';
import { initReactI18next } from 'react-i18next';
const moment = require('moment');



const defaultLanguage = 'en';

i18n.use(initReactI18next).init({
  lng: defaultLanguage,
  fallbackLng: 'en',
  defaultNS: 'form',
  resources: {
    en,
  },
  interpolation: {
    function(value, format, lng) {
      if (value instanceof Date) return moment(value).format(format);
      return value;
    },
  },
  react: {
    wait: false,
    bindI18n: false,
    bindStore: false,
    nsMode: false,
  },
});

i18n.on('languageChanged', currentLang => {
  moment.locale(currentLang);
});

moment.locale(defaultLanguage);

export default i18n;
