import './App.css';

import { useEffect, useState, useContext } from "react";
import { Routes, Route, useNavigate } from "react-router-dom";
import { useMutation, useQuery } from 'react-query';
import { UserContext } from './context/userContext';

import { API, setAuthToken } from './config/api';

import Landing from "./pages/Landing";
import PostsPage from "./pages/PostsPage";
import PostDetails from "./pages/PostDetails";
import Profile from "./pages/Profile";
import UpdateProfile from "./pages/UpdateProfile";
import UploadPost from "./pages/UploadPost";
import SendProject from "./pages/SendProject";
import UploadArt from "./pages/UploadArt";
import Order from "./pages/Order";
import Project from "./pages/Project";
import OrderList from "./pages/OrderList";


import PrivateRoute from "./components/PrivateRoute";

import NavbarSection from "./components/NavbarSection";
import LoginModal from "./components/LoginModal";
import RegisterModal from "./components/RegisterModal";
import SuccessRegisterModal from "./components/SuccessRegisterModal";
import UnregisteredEmailModal from "./components/UnregisteredEmailModal";
import RegisteredEmailModal from "./components/RegisteredEmailModal";
import WrongPasswordModal from "./components/WrongPasswordModal";
import SuccessLoginToast from "./components/SuccessLoginToast";
import SuccessLogoutToast from "./components/SuccessLogoutToast";
import SuccessAddArtModal from "./components/SuccessAddArtModal";
import SuccessUpdateProfileModal from "./components/SuccessUpdateProfileModal";
import SuccessUploadPostModal from "./components/SuccessUploadPostModal";
import SuccessMakeOrderModal from "./components/SuccessMakeOrderModal";
import FailedMakeOrderModal from "./components/FailedMakeOrderModal";
import SuccessUpdateOrderStatusModal from "./components/SuccessUpdateOrderStatusModal";
import ImageEnlargerModal from './components/ImageEnlargerModal';

function App() {
  const navigate = useNavigate();
  useEffect(() => window.scroll({top: 0, behavior: "smooth"}),[]);
  const [state, dispatch] = useContext(UserContext);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    if (!isLoading) {
      if (state.isLogin === false) {
        navigate('/');
      }
    }
  }, [isLoading]);

  useEffect(() => {
    if (localStorage.token) {
      setAuthToken(localStorage.token);
      checkUser();
    } else {
      setTimeout(() => {
        setIsLoading(false)
      }, 3000);
    }
  }, []);

  const checkUser = async () => {
    try {
      const response = await API.get('/check-auth');
      let payload = response.data.data;
      payload.token = localStorage.token;
      dispatch({
        type: 'USER_SUCCESS',
        payload,
      });
      setTimeout(() => {
        setIsLoading(false)
      }, 3000);
    } catch (error) {
      dispatch({
        type: 'AUTH_ERROR',
      });
      setTimeout(() => {
        setIsLoading(false)
      }, 3000);
      return
    }
  };

  const [Posts, SetPosts] = useState([]);
  const [Followings, SetFollowings] = useState([]);
  const [Profiles, SetProfiles] = useState([]);
  const [Orders, SetOrders] = useState([]);

  useQuery('postsCache', async () => {
    const response = await API.get('/posts');
    SetPosts(response.data.data);
  });
  useQuery('followingsCache', async () => {
    const response = await API.get('/followings');
    SetFollowings(response.data.data);
  });
  useQuery('profilesCache', async () => {
    const response = await API.get('/profiles');
    SetProfiles(response.data.data);
  });
  useQuery('ordersCache', async () => {
    const response = await API.get('/orders');
    SetOrders(response.data.data);
  });

  const [formLogin, setFormLogin] = useState({
    email: "",
    password: ""
  });
  const formLoginHandleOnChange = (e) => {
    setFormLogin({
      ...formLogin,
      [e.target.name]: e.target.value,
    });
  };
  const formLoginHandleOnSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();
      
      const responseLogin = await API.post('/login', formLogin);
      setAuthToken(responseLogin.data.data.token);

      const responseUser = await API.get('/user/' + responseLogin.data.data.id);
      responseUser.data.data.token = responseLogin.data.data.token;
      
      dispatch({
        type: 'LOGIN_SUCCESS',
        payload: responseUser.data.data,
      });
      settoastSuccessLogin(true);
      navigate("/posts");

      setFormLogin((formLogin) => ({
        ...formLogin,
        email: "",
        password: ""
      }));
  
      setModalLoginShow(false);
    } catch (error) {
      if (error.response.data.message === "record not found")  setModalUnregisteredEmail(true);
      else if (error.response.data.message === "wrong password") setModalWrongPassword(true);
      setFormLogin((formLogin) => ({
        ...formLogin,
        email: "",
        password: ""
      }));
  
      setModalLoginShow(false);
    }
  });

  const [formRegister, setFormRegister] = useState({
    name: "",
    email: "",
    password: "",
  });
  const formRegisterHandleOnChange = (e) => {
    setFormRegister({
      ...formRegister,
      [e.target.name]: e.target.value,
    });
  };
  const formRegisterHandleOnSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();
      await API.post('/register', formRegister);

      setFormRegister((formRegister) => ({
        ...formRegister,
        name: "",
        email: "",
        password: "",
      }));
      setModalRegisterShow(false);
      setmodalSuccessRegister(true);
    } catch (error) {
      if (error.response.data.message === "This email is already registered") setModalRegisteredEmail(true);
      setFormRegister((formRegister) => ({
        ...formRegister,
        name: "",
        email: "",
        password: "",
      }));

      setModalRegisterShow(false);
    }
  });

  const [modalLoginShow, setModalLoginShow] = useState(false);
  const [modalRegisterShow, setModalRegisterShow] = useState(false);
  const [modalSuccessRegister, setmodalSuccessRegister] = useState(false);
  const [modalUnregisteredEmail, setModalUnregisteredEmail] = useState(false);
  const [modalRegisteredEmail, setModalRegisteredEmail] = useState(false);
  const [modalWrongPassword, setModalWrongPassword] = useState(false);
  const [toastSuccessLogin, settoastSuccessLogin] = useState(false);
  const [toastSuccessLogout, settoastSuccessLogout] = useState(false);
  const [modalAddArt, setModalAddArt] = useState(false);
  const [modalUpdateProfile, setModalUpdateProfile] = useState(false);
  const [modalUploadPost, setModalUploadPost] = useState(false);
  const [modalMakeOrder, setModalMakeOrder] = useState(false);
  const [modalFailedMakeOrder, setModalFailedMakeOrder] = useState(false);
  const [modalImageEnlarger, setModalImageEnlarger] = useState(false);
  const [showSuccessUpdateOrderStatusModal, setShowSuccessUpdateOrderStatusModal] = useState(false);
  const [message, setMessage] = useState("");
  const [imageToEnlarge, setImageToEnlarge] = useState("");

  const handleShowModalImageEnlarger = (image) => {
    setImageToEnlarge(image);
    setModalImageEnlarger(true);
  }

  const [navbarOn, setNavbarOn] = useState(false);

  return (
    <>
      {isLoading ? (
        <div className="tw-flex tw-justify-center tw-items-center" style={{width:"100vw",height:"100vh"}}>
          <img src="/images/title.webp" alt="WaysGallery" className="tw-animate-ping tw-absolute"/>
          <img src="/images/title.webp" alt="WaysGallery" className="tw-relative"/>
        </div>
      ) : 
        <>
          <ImageEnlargerModal 
            show={modalImageEnlarger} 
            onHide={() => setModalImageEnlarger(false)} 
            imageToEnlarge={imageToEnlarge}
          />
          <SuccessUpdateOrderStatusModal 
            show={showSuccessUpdateOrderStatusModal} 
            onHide={() => setShowSuccessUpdateOrderStatusModal(false)}
            message={message}
          />
          <FailedMakeOrderModal 
            show={modalFailedMakeOrder} 
            onHide={() => setModalFailedMakeOrder(false)} 
          />
          <SuccessMakeOrderModal 
            show={modalMakeOrder} 
            onHide={() => setModalMakeOrder(false)} 
          />
          <SuccessUploadPostModal 
            show={modalUploadPost} 
            onHide={() => setModalUploadPost(false)} 
          />
          <SuccessUpdateProfileModal 
            show={modalUpdateProfile} 
            onHide={() => setModalUpdateProfile(false)} 
          />
          <SuccessAddArtModal 
            show={modalAddArt} 
            onHide={() => setModalAddArt(false)} 
          />
          <SuccessLoginToast 
            show={toastSuccessLogin} 
            onClose={() => settoastSuccessLogin(false)} 
          />
          <SuccessLogoutToast 
            show={toastSuccessLogout} 
            onClose={() => settoastSuccessLogout(false)} 
          />
          <SuccessRegisterModal  
              show={modalSuccessRegister} 
              onHide={() => {
                setmodalSuccessRegister(false);
                setModalLoginShow(true);
              }} 
          />
          <UnregisteredEmailModal 
            show={modalUnregisteredEmail} 
            onHide={() => {
              setModalUnregisteredEmail(false);
              setModalRegisterShow(true);
            }} 
          />
          <RegisteredEmailModal 
            show={modalRegisteredEmail} 
            onHide={() => {
              setModalRegisteredEmail(false);
              setModalLoginShow(true);
            }} 
          />
          <WrongPasswordModal 
            show={modalWrongPassword} 
            onHide={() => {
              setModalWrongPassword(false);
              setModalLoginShow(true);
            }} 
          />
          <LoginModal 
            show={modalLoginShow} 
            onHide={() => {
              setModalLoginShow(false);
              setFormLogin((formLogin) => ({
                ...formLogin,
                email: "",
                password: "",
              }));
            }} 
            changeModal={() => {
              setModalLoginShow(false);
              setModalRegisterShow(true);
              setFormLogin((formLogin) => ({
                ...formLogin,
                email: "",
                password: "",
              }));
            }} 
            formLogin={formLogin} 
            loginOnChange={(e) => formLoginHandleOnChange(e)}
            loginOnSubmit={(e) => formLoginHandleOnSubmit.mutate(e)}
          />
          <RegisterModal 
            show={modalRegisterShow} 
            onHide={() => {
              setModalRegisterShow(false);
              setFormRegister((formRegister) => ({
                ...formRegister,
                email: "",
                password: "",
                name: "",
              }));
            }} 
            changeModal={() => {
              setModalRegisterShow(false);
              setModalLoginShow(true);
              setFormRegister((formRegister) => ({
                ...formRegister,
                email: "",
                password: "",
                name: "",
              }));
            }}
            formRegister={formRegister} 
            registerOnChange={(e) => formRegisterHandleOnChange(e)}
            registerOnSubmit={(e) => formRegisterHandleOnSubmit.mutate(e)}
          />

          {navbarOn && 
            <>
              <svg onClick={() => window.scroll({top: 0, behavior: "smooth"})} xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" className="position-fixed" style={{ bottom:"1rem", right: "1rem", width:"2.5rem", cursor:"pointer", zIndex:"5", filter: `drop-shadow(1px 1px 0 #E7E7E7) drop-shadow(-1px -1px 0 #E7E7E7) drop-shadow(1px -1px 0 #E7E7E7) drop-shadow(-1px 1px 0 #E7E7E7)`, backgroundColor:`#E7E7E7`, borderRadius:"50%" }}>
                <path fill="#2FC4B2" d="M256 512A256 256 0 1 0 256 0a256 256 0 1 0 0 512zM135.1 217.4l107.1-99.9c3.8-3.5 8.7-5.5 13.8-5.5s10.1 2 13.8 5.5l107.1 99.9c4.5 4.2 7.1 10.1 7.1 16.3c0 12.3-10 22.3-22.3 22.3H304v96c0 17.7-14.3 32-32 32H240c-17.7 0-32-14.3-32-32V256H150.3C138 256 128 246 128 233.7c0-6.2 2.6-12.1 7.1-16.3z"/>
              </svg>
              <NavbarSection 
                LoggedInUserId={state.user.id} 
                User={state.user} 
                Profiles={Profiles} 
                showModalLogin={() => setModalLoginShow(true)} 
                showModalRegister={() => setModalRegisterShow(true)} 
                isLogin={state.isLogin} 
                logout={() => {
                  dispatch({
                    type: "LOGOUT"
                  });
                  settoastSuccessLogout(true);
                  navigate("/");
                  setNavbarOn(false);
                }} 
              />
            </>
          }
          <Routes>
            <Route path="/" element={ <Landing 
              isLogin={state.isLogin} 
              setNavbarOff={() => setNavbarOn(false)}
              showModalLogin={() => setModalLoginShow(true)} 
              showModalRegister={() => setModalRegisterShow(true)} 
            /> } />
            <Route path="/posts" element={ <PostsPage 
              isLogin={state.isLogin}
              User={state.user} 
              Followings={Followings} 
              Posts={Posts} 
              showModalLogin={() => setModalLoginShow(true)} 
              setNavbarOn={() => setNavbarOn(true)} 
            /> } />
            <Route path="/posts/:id" element={ <PostDetails 
              isLogin={state.isLogin} 
              User={state.user} 
              Profiles={Profiles}
              Followings={Followings} 
              SetFollowings={SetFollowings} 
              Posts={Posts} 
              showModalLogin={() => setModalLoginShow(true)} 
              setNavbarOn={() => setNavbarOn(true)} 
              handleShowModalImageEnlarger={handleShowModalImageEnlarger} 
            /> } />
            <Route path="/profile/:id" element={ <Profile  
              isLogin={state.isLogin} 
              setNavbarOn={() => setNavbarOn(true)} 
              LoggedInUserId={state.user.id} 
              User={state.user} 
              Profiles={Profiles} 
              Followings={Followings} 
              SetFollowings={SetFollowings} 
              Posts={Posts} 
              showModalLogin={() => setModalLoginShow(true)} 
              handleShowModalImageEnlarger={handleShowModalImageEnlarger} 
            /> } />
            <Route path="/" element={<PrivateRoute />}>
              <Route path="/update-profile" element={ <UpdateProfile  
                LoggedInUserId={state.user.id} 
                User={state.user} 
                setNavbarOn={() => setNavbarOn(true)} 
                Profiles={Profiles} 
                SetProfiles={SetProfiles} 
                showModalSuccessUpdateProfile={() => setModalUpdateProfile(true)} 
              /> } />
              <Route path="/upload-post" element={ <UploadPost  
                LoggedInUserId={state.user.id} 
                setNavbarOn={() => setNavbarOn(true)} 
                Posts={Posts} 
                SetPosts={SetPosts} 
                showModalSuccessUploadPost={() => setModalUploadPost(true)}
              /> } />
              <Route path="/send-project/:id" element={ <SendProject  
                setNavbarOn={() => setNavbarOn(true)} 
                SetOrders={SetOrders} 
                setShowSuccessUpdateOrderStatusModal={() => setShowSuccessUpdateOrderStatusModal(true)} 
                setMessage={() => setMessage("You have successfully sent the project of this order. Thank you.")}
              /> } />
              <Route path="/upload-art" element={ <UploadArt  
                setNavbarOn={() => setNavbarOn(true)} 
                LoggedInUserId={state.user.id} 
                showModalSuccessAddArt={() => setModalAddArt(true)} 
                SetProfiles={SetProfiles}
              /> } />
              <Route path="/order/:id" element={ <Order  
                setNavbarOn={() => setNavbarOn(true)} 
                Orders={Orders} 
                SetOrders={SetOrders} 
                showModalSuccessMakeOrder={() => setModalMakeOrder(true)} 
                showModalFailedMakeOrder={() => setModalFailedMakeOrder(true)} 
              /> } />
              <Route path="/project/:id" element={ <Project  
                setNavbarOn={() => setNavbarOn(true)} 
                Orders={Orders} 
                handleShowModalImageEnlarger={handleShowModalImageEnlarger} 
              /> } />
              <Route path="/order-list" element={ <OrderList  
                setNavbarOn={() => setNavbarOn(true)} 
                LoggedInUserId={state.user.id} 
                Profiles={Profiles}
                Orders={Orders} 
                SetOrders={SetOrders} 
                setShowSuccessUpdateOrderStatusSuccessModal={() => setShowSuccessUpdateOrderStatusModal(true)} 
                setMessageSuccess={() => setMessage("You have successfully approved this order. Thank you.")} 
                setShowSuccessUpdateOrderStatusCancelModal={() => setShowSuccessUpdateOrderStatusModal(true)} 
                setMessageCancel={() => setMessage("You have successfully canceled this order. Thank you.")} 
              /> } />
            </Route>
          </Routes>
        </>
      }
    </>
  );
}

export default App;
