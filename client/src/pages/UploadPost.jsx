import { Container, Form, Button } from "react-bootstrap";
import { useState } from "react";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";
import { API } from '../config/api';

export default function UploadPost(props) {
  document.title = "Upload Post | WaysGallery";
  props.setNavbarOn();
  const navigate = useNavigate();

  const [image1, setImage1] = useState("");
  const [image2, setImage2] = useState("");
  const [image3, setImage3] = useState("");
  const [image4, setImage4] = useState("");
  const [image5, setImage5] = useState("");
  const [formPost, setFormPost] = useState({
    title: "",
    description: "",
    image1: "",
    image2: "",
    image3: "",
    image4: "",
    image5: "",
  });
  const formPostHandleOnChange = (e) => {
    setFormPost({
      ...formPost,
      [e.target.name]: e.target.value,
    });
  };
  const formPostHandleOnChangeImage1 = (e) => {
    setFormPost({
      ...formPost,
      [e.target.name]:
        e.target.type === 'file' ? e.target.files : e.target.value,
    });
    if (e.target.type === 'file') {
      let url = URL.createObjectURL(e.target.files[0]);
      setImage1(url);
    }
  };
  const formPostHandleOnChangeImage2 = (e) => {
    setFormPost({
      ...formPost,
      [e.target.name]:
        e.target.type === 'file' ? e.target.files : e.target.value,
    });
    if (e.target.type === 'file') {
      let url = URL.createObjectURL(e.target.files[0]);
      setImage2(url);
    }
  };
  const formPostHandleOnChangeImage3 = (e) => {
    setFormPost({
      ...formPost,
      [e.target.name]:
        e.target.type === 'file' ? e.target.files : e.target.value,
    });
    if (e.target.type === 'file') {
      let url = URL.createObjectURL(e.target.files[0]);
      setImage3(url);
    }
  };
  const formPostHandleOnChangeImage4 = (e) => {
    setFormPost({
      ...formPost,
      [e.target.name]:
        e.target.type === 'file' ? e.target.files : e.target.value,
    });
    if (e.target.type === 'file') {
      let url = URL.createObjectURL(e.target.files[0]);
      setImage4(url);
    }
  };
  const formPostHandleOnChangeImage5 = (e) => {
    setFormPost({
      ...formPost,
      [e.target.name]:
        e.target.type === 'file' ? e.target.files : e.target.value,
    });
    if (e.target.type === 'file') {
      let url = URL.createObjectURL(e.target.files[0]);
      setImage5(url);
    }
  };
  const formPostHandleOnSubmit = useMutation(async (e) => {
      e.preventDefault();
  
      const config = {
        headers: {
          "Content-type": "multipart/form-data",
        },
      };

      const formData = new FormData();
      if (formPost.image1) {
        formData.set('image1', formPost?.image1[0], formPost?.image1[0]?.name);
      }
      if (formPost.image2) {
        formData.set('image2', formPost?.image2[0], formPost?.image2[0]?.name);
      }
      if (formPost.image3) {
        formData.set('image3', formPost?.image3[0], formPost?.image3[0]?.name);
      }
      if (formPost.image4) {
        formData.set('image4', formPost?.image4[0], formPost?.image4[0]?.name);
      }
      if (formPost.image5) {
        formData.set('image5', formPost?.image5[0], formPost?.image5[0]?.name);
      }
      formData.set('title', formPost.title);
      formData.set('description', formPost.description);

      try {
        const response = await API.post('/post', formData, config);

        const newPost = {
          id: response.data.data.id,
          name: response.data.data.name,
          description: response.data.data.description,
          image_1: response.data.data.image_1,
          image_2: response.data.data.image_2,
          image_3: response.data.data.image_3,
          image_4: response.data.data.image_4,
          image_5: response.data.data.image_5,
          user: {id: props.LoggedInUserId},
        }
        props.SetPosts([...props.Posts, newPost]);

        setFormPost((formPost) => ({
          ...formPost,
          title: "",
          description: "",
          image1: "",
          image2: "",
          image3: "",
          image4: "",
          image5: "",
        }));
        setImage1("");
        setImage2("");
        setImage3("");
        setImage4("");
        setImage5("");
        props.showModalSuccessUploadPost();
        navigate("/posts");
      }
      catch (error) {
        return
      }
  });

  return (
    <Container className="tw-min-h-full">
      <Form onSubmit={(e) => formPostHandleOnSubmit.mutate(e)} className="md:tw-flex md:tw-justify-between tw-pb-20">
        <div className="animate__animated animate__slideInLeft tw-w-full md:tw-w-6/12 tw-mt-28 md:tw-mt-0">
          {image1 === "" ? (
              <label for="image1" className="tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-5 tw-border-dashed tw-rounded tw-flex tw-flex-col tw-justify-center tw-items-center tw-mt-28 tw-cursor-pointer">
                <img src="/images/icon-cloud-upload.png" alt="Post Preview 1" className="tw-object-contain"/>
                <p><strong className="tw-text-custom-primary">Browse</strong> to choose a file</p>
              </label>
            ) : (
              <label for="image1" className="tw-flex tw-flex-col tw-justify-center tw-items-center tw-mt-28 tw-cursor-pointer tw-rounded">
                <img src={image1} alt="Post Preview 1" className="tw-w-full tw-object-cover"/>
              </label>
            )
          }
          <div className="tw-flex tw-justify-between tw-mt-5">
            {image2 === "" ? (
              <label for="image2" className="tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-2 md:tw-p-5 tw-border-dashed tw-rounded tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-mr-4 tw-cursor-pointer">
                <img src="/images/icon-add.png" alt="Post Preview 1" className="tw-object-contain"/>
              </label>
            ) : (
              <label for="image2" className="tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-mr-4 tw-cursor-pointer tw-rounded">
                <img src={image2} alt="Post Preview 1" className="tw-h-full tw-object-cover"/>
              </label>
            )}
            {image3 === "" ? (
              <label for="image3" className="tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-2 md:tw-p-5 tw-border-dashed tw-rounded tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-mr-4 tw-cursor-pointer">
                <img src="/images/icon-add.png" alt="Post Preview 1" className="tw-object-contain"/>
              </label>
            ) : (
              <label for="image3" className="tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-mr-4 tw-cursor-pointer tw-rounded">
                <img src={image3} alt="Post Preview 1" className="tw-h-full tw-object-cover"/>
              </label>
            )}
            {image4 === "" ? (
              <label for="image4" className="tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-2 md:tw-p-5 tw-border-dashed tw-rounded tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-mr-4 tw-cursor-pointer">
                <img src="/images/icon-add.png" alt="Post Preview 1" className="tw-object-contain"/>
              </label>
            ) : (
              <label for="image4" className="tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-mr-4 tw-cursor-pointer tw-rounded">
                <img src={image4} alt="Post Preview 1" className="tw-h-full tw-object-cover"/>
              </label>
            )}
            {image5 === "" ? (
              <label for="image5" className="tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-2 md:tw-p-5 tw-border-dashed tw-rounded tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-cursor-pointer">
                <img src="/images/icon-add.png" alt="Post Preview 1" className="tw-object-contain"/>
              </label>
            ) : (
              <label for="image5" className="tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-cursor-pointer">
                <img src={image5} alt="Post Preview 1" className="tw-h-full tw-object-cover"/>
              </label>
            )}
          </div>
          <input id="image1" name="image1" onChange={(e) => formPostHandleOnChangeImage1(e)} type="file" className="tw-hidden"/>
          <input id="image2" name="image2" onChange={(e) => formPostHandleOnChangeImage2(e)} type="file" className="tw-hidden"/>
          <input id="image3" name="image3" onChange={(e) => formPostHandleOnChangeImage3(e)} type="file" className="tw-hidden"/>
          <input id="image4" name="image4" onChange={(e) => formPostHandleOnChangeImage4(e)} type="file" className="tw-hidden"/>
          <input id="image5" name="image5" onChange={(e) => formPostHandleOnChangeImage5(e)} type="file" className="tw-hidden"/>
        </div>
        <div className="animate__animated animate__slideInRight md:!tw-px-4 !tw-pb-4 tw-flex tw-flex-col tw-items-center tw-mt-10 tw-w-full md:tw-w-5/12 md:tw-mt-28">
          <Form.Group className="!tw-mb-6 tw-w-full" controlId="formGreeting">
            <Form.Control type="text" name="title" onChange={(e) => formPostHandleOnChange(e)} value={formPost.title} placeholder="Title" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" required/>
          </Form.Group>
          <Form.Group className="!tw-mb-6 tw-w-full" controlId="formDescription">
            <Form.Control as="textarea" onChange={(e) => formPostHandleOnChange(e)} placeholder="Description" value={formPost.description} name="description" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" rows={8} required/>
          </Form.Group>
          <div className="tw-flex tw-justify-center tw-w-full">
            <Button onClick={() => navigate(`/profile/${props.LoggedInUserId}`)} className="!tw-bg-custom-secondary hover:!tw-bg-custom-secondary-dark !tw-text-black !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32 tw-mr-4">Cancel</Button>
            <Button type="submit" className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32">Save</Button>
          </div>
        </div>
      </Form>
    </Container>
  );
}