import { Container, Form, Button } from "react-bootstrap";
import { useState, useEffect } from "react";
import { useMutation } from "react-query";
import { useNavigate, useParams } from "react-router-dom";
import { API } from '../config/api';

export default function Order(props) {
  document.title = "Hire | WaysGallery";
  props.setNavbarOn();
  const navigate = useNavigate();
  let params = useParams();

  useEffect(() => {
    const midtransScriptUrl = "https://app.sandbox.midtrans.com/snap/snap.js";
    const myMidtransClientKey = process.env.REACT_APP_MIDTRANS_CLIENT_KEY;
  
    let scriptTag = document.createElement("script");
    scriptTag.src = midtransScriptUrl;
    scriptTag.setAttribute("data-client-key", myMidtransClientKey);
  
    document.body.appendChild(scriptTag);
    return () => {
      document.body.removeChild(scriptTag);
    };
  }, []);

  const [formOrder, setFormOrder] = useState({
    title: "",
    description: "",
    start_date: "",
    end_date: "",
    price: 0,
  });
  const formOrderHandleOnChange = (e) => {
    setFormOrder({
      ...formOrder,
      [e.target.name]: e.target.value,
    });
  };
  const formOrderHandleOnSubmit = useMutation(async (e) => {
    e.preventDefault();

    const formData = new FormData();
    formData.set('title', formOrder.title);
    formData.set('description', formOrder.description);
    formData.set('start_date', formOrder.start_date);
    formData.set('end_date', formOrder.end_date);
    formData.set('price', formOrder.price);

    try {
      const response = await API.post(`/order/${params.id}`, formData);
      const token = response.data.data.token;

      window.snap.pay(token, {
        onSuccess: function (result) {
          const newOrder = {
            id: response.data.data.id,
            name: response.data.data.name,
            description: response.data.data.description,
            start_date: response.data.data.start_date,
            end_date: response.data.data.end_date,
            price: response.data.data.price,
          }
          props.SetOrders([...props.Orders, newOrder]);
    
          setFormOrder((formOrder) => ({
            ...formOrder,
            title: "",
            description: "",
            start_date: "",
            end_date: "",
            price: 0,
          }));
          props.showModalSuccessMakeOrder();
          navigate("/order-list");
        },
        onPending: function (result) {
          const newOrder = {
            id: response.data.data.id,
            name: response.data.data.name,
            description: response.data.data.description,
            start_date: response.data.data.start_date,
            end_date: response.data.data.end_date,
            price: response.data.data.price,
          }
          props.SetOrders([...props.Orders, newOrder]);
    
          setFormOrder((formOrder) => ({
            ...formOrder,
            title: "",
            description: "",
            start_date: "",
            end_date: "",
            price: 0,
          }));
          props.showModalSuccessMakeOrder();
          navigate("/order-list");
        },
        onError: function (result) {
          props.showModalFailedMakeOrder();
          navigate("/order-list");
        },
        onClose: function () {
          props.showModalFailedMakeOrder();
          navigate("/order-list");
        },
      });
    }
    catch (error) {
      return
    }
});

  return (
    <Container className="tw-min-h-full tw-flex tw-justify-center">
      <Form onSubmit={(e) => formOrderHandleOnSubmit.mutate(e)} className="animate__animated animate__zoomIn tw-mt-28 tw-w-7/12">
        <Form.Group className="!tw-mb-6 tw-w-full" controlId="formGreeting">
          <Form.Control type="text" name="title" onChange={(e) => formOrderHandleOnChange(e)} value={formOrder.title} placeholder="Title" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" required/>
        </Form.Group>
        <Form.Group className="!tw-mb-6 tw-w-full" controlId="formDescription">
          <Form.Control as="textarea" placeholder="Description" name="description" onChange={(e) => formOrderHandleOnChange(e)} value={formOrder.description} className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" rows={4} required/>
        </Form.Group>
        <Form.Group className="!tw-mb-6 tw-w-full !tw-flex" controlId="formPrice">
          <Form.Control type="date" name="start_date" onChange={(e) => formOrderHandleOnChange(e)} value={formOrder.start_date} placeholder="Start Date" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary !tw-w-6/12 !tw-mr-4" required/>
          <Form.Control type="date" name="end_date" onChange={(e) => formOrderHandleOnChange(e)} value={formOrder.end_date} placeholder="End Date" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary !tw-w-6/12" required/>
        </Form.Group>
        <Form.Group className="!tw-mb-6 tw-w-full" controlId="formPrice">
          <Form.Control type="number" name="price" onChange={(e) => formOrderHandleOnChange(e)} value={formOrder.price} placeholder="Price" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" required/>
        </Form.Group>
        <div className="tw-flex tw-justify-center tw-w-full tw-mt-16">
          <Button onClick={() => navigate(`/profile/${parseInt(params.id)}`)} className="!tw-bg-custom-secondary hover:!tw-bg-custom-secondary-dark !tw-text-black !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32 tw-mr-4">Cancel</Button>
          <Button type="submit" className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32">Bidding</Button>
        </div>
      </Form>
    </Container>
  );
}