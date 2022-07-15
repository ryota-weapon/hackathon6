// Import the functions you need from the SDKs you need
import { initializeApp,  } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import { getAuth } from "firebase/auth";

// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyCclJqyGUJ0EYSWvIhNGzYsy-qafIIl5fA",
  authDomain: "hackathon6-64cf3.firebaseapp.com",
  projectId: "hackathon6-64cf3",
  storageBucket: "hackathon6-64cf3.appspot.com",
  messagingSenderId: "576726558986",
  appId: "1:576726558986:web:96c3c69f0a1cb2837572f3",
  measurementId: "G-SKJ7CSY1DL"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);
// console.log(analytics);


export const auth = getAuth();
