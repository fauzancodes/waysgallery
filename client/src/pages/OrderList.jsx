import { Container, Dropdown, Button, Table } from "react-bootstrap";
import { useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { API } from '../config/api';

import DetailOrderModal from "../components/DetailOrderModal";

function OrderList(props) {
  document.title = "Orders | WaysGallery";
  const navigate = useNavigate();
  props.setNavbarOn();

  const Orders = [...props.Orders];
  Orders.sort((a, b) => b.id - a.id);

  const [orderFilter, setOrderFilter] = useState("order");
  useEffect(() => {
    const postFilterStorage = localStorage.getItem("WaysGalleryOrderFilter");
    if (postFilterStorage === "offer") setOrderFilter("offer");
    else {
      localStorage.setItem("WaysGalleryOrderFilter", "order");
      setOrderFilter("order")
    };
  }, []);
  const toggleOrderFilter = () => {
    if (orderFilter === "order") {
      setOrderFilter("offer");
      localStorage.setItem("WaysGalleryOrderFilter", "offer");
    }
    else {
      setOrderFilter("order");
      localStorage.setItem("WaysGalleryOrderFilter", "order");
    }
  }

  const [showDetailOrder, setShowDetailOrder] = useState(false);
  const [orderTitle, setOrderTitle] = useState("");
  const [orderDescription, setOrderDescription] = useState("");
  const [orderPrice, setOrderPrice] = useState(0);

  const handleDetailOrderModal = (title, description, price) => {
    setOrderTitle(title);
    setOrderDescription(description);
    setOrderPrice(price);
    setShowDetailOrder(true);
  }

  const handleApprove = async (id) => {
    try {
      await API.patch('/order/' + id, {"status":"success"});
      const newOrderStatus = {
        status: "success",
      }
      props.SetOrders(prevOrders => {
        return prevOrders.map(order => {
          if (order.id === id) {
            return { ...order, ...newOrderStatus };
          }
          return order;
        });
      });
      props.setMessageSuccess();
      props.setShowSuccessUpdateOrderStatusSuccessModal();
    }
    catch (error) {
      return
    }
  }
  const handleCancel = async (id) => {
    try {
      await API.patch('/order/' + id, {"status":"cancel"});
      const newOrderStatus = {
        status: "cancel",
      }
      props.SetOrders(prevOrders => {
        return prevOrders.map(order => {
          if (order.id === id) {
            return { ...order, ...newOrderStatus };
          }
          return order;
        });
      });
      props.setMessageCancel()
      props.setShowSuccessUpdateOrderStatusCancelModal();
    }
    catch (error) {
      return
    }
  }
  const handleSend = async (id) => {
    navigate("/send-project/" + id)
  }
  const handleView = (id) => {
    navigate("/project/" + id)
  }

  const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];

  return (
    <>
      <DetailOrderModal 
        show={showDetailOrder} 
        onHide={() => setShowDetailOrder(false)} 
        title={orderTitle} 
        description={orderDescription} 
        price={orderPrice} 
      />
      <Container className="tw-min-h-full">
        <Dropdown className="tw-order-2 md:tw-order-1 tw-mt-28 tw-mb-10">
          <Dropdown.Toggle className="tw-w-40 !tw-bg-custom-secondary !tw-border-0 hover:!tw-bg-custom-secondary-dark !tw-text-black !tw-font-bold">
            {orderFilter === "order" ? "Order" : "Offer"}
          </Dropdown.Toggle>
          <Dropdown.Menu className="!tw-py-0">
            <Dropdown.Item onClick={toggleOrderFilter}>{orderFilter === "order" ? "Offer" : "Order"}</Dropdown.Item>
          </Dropdown.Menu>
        </Dropdown>
        {orderFilter === "order" ? (
          Orders.filter(order => order.client_id === props.LoggedInUserId).length > 0 ? (
            <Table striped bordered hover className="animate__animated animate__zoomIn">
            <thead>
              <tr>
                <th>No</th>
                <th>Vendor</th>
                <th>Order</th>
                <th>Start Date</th>
                <th>End Date</th>
                <th>Status</th>
                <th>Message</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              {
                Orders.filter(order => order.client_id === props.LoggedInUserId).map((order, index) => (
                  <tr key={index}>
                    <td>{index+1}</td>
                    <td>{props.Profiles.find(profile => profile.id === order.vendor_id).name}</td>
                    <td onClick={() => handleDetailOrderModal(order.title, order.description, order.price)} className="tw-cursor-pointer !tw-text-custom-primary tw-underline">{order.title}</td>
                    <td>{`${new Date(order.start_date.split("T")[0]).getDate()} ${months[new Date(order.start_date.split("T")[0]).getMonth()]} ${new Date(order.start_date.split("T")[0]).getFullYear()}`}</td>
                    <td>{`${new Date(order.end_date.split("T")[0]).getDate()} ${months[new Date(order.end_date.split("T")[0]).getMonth()]} ${new Date(order.end_date.split("T")[0]).getFullYear()}`}</td>
                    <td className={
                      order.status === "waiting" ? "!tw-text-custom-warning" : 
                      order.status === "success" ? "!tw-text-custom-success" : 
                      order.status === "cancel" ? "!tw-text-custom-danger" : 
                      order.status === "complete" ? "!tw-text-custom-primary" : null
                    }>{order.status}</td>
                    <td>{
                      order.status === "waiting" ? "Please wait, your vendor will soon approve your order." : 
                      order.status === "success" ? "Your vendor has agreed to take your order, please wait until the order is completed by your vendor." : 
                      order.status === "cancel" ? "Sorry, your vendor did not approve your order. Please find another vendor." : 
                      order.status === "complete" ? `This order has been completed. If you want to contact your vendor, please contact via ${props.Profiles.find(profile => profile.id === order.vendor_id).user.email}` : null
                    }</td>
                    <td className="tw-align-middle">
                      {
                        order.status === "waiting" ? <img src="/images/icon-loading.png" alt={order.status} className="tw-mx-auto" /> : 
                        order.status === "success" ? <img src="/images/icon-success.png" alt={order.status} className="tw-mx-auto" /> : 
                        order.status === "cancel" ? <img src="/images/icon-cancel.png" alt={order.status} className="tw-mx-auto" /> : 
                        order.status === "complete" ? <Button onClick={() => handleView(order.id)} className="!tw-bg-custom-success hover:!tw-bg-custom-success-dark !tw-border-0">View Project</Button> : null
                      }
                    </td>
                  </tr>
                ))
              }
            </tbody>
          </Table>  
          ) : (
            <p className="tw-opacity-50">There are no orders to display.</p>
          )
        ) : (
          Orders.filter(order => order.vendor_id === props.LoggedInUserId).length > 0 ? (
            <Table striped bordered hover className="animate__animated animate__zoomIn">
              <thead>
                <tr>
                  <th>No</th>
                  <th>Client</th>
                  <th>Order</th>
                  <th>Start Date</th>
                  <th>End Date</th>
                  <th>Status</th>
                  <th>Message</th>
                  <th>Action</th>
                </tr>
              </thead>
              <tbody>
              {
                Orders.filter(order => order.vendor_id === props.LoggedInUserId).map((order, index) => (
                  <tr key={index}>
                    <td>{index+1}</td>
                    <td>{order.user.name}</td>
                    <td onClick={() => handleDetailOrderModal(order.title, order.description, order.price)} className="tw-cursor-pointer !tw-text-custom-primary tw-underline">{order.title}</td>
                    <td>{`${new Date(order.start_date.split("T")[0]).getDate()} ${months[new Date(order.start_date.split("T")[0]).getMonth()]} ${new Date(order.start_date.split("T")[0]).getFullYear()}`}</td>
                    <td>{`${new Date(order.end_date.split("T")[0]).getDate()} ${months[new Date(order.end_date.split("T")[0]).getMonth()]} ${new Date(order.end_date.split("T")[0]).getFullYear()}`}</td>
                    <td className={
                      order.status === "waiting" ? "!tw-text-custom-warning" : 
                      order.status === "success" ? "!tw-text-custom-success" : 
                      order.status === "cancel" ? "!tw-text-custom-danger" : 
                      order.status === "complete" ? "!tw-text-custom-primary" : null
                    }>{order.status}</td>
                    <td>{
                      order.status === "waiting" ? "Please check the details of this order and agree if you wish." : 
                      order.status === "success" ? "You have agreed to this order. Please give your best so that your client is satisfied." : 
                      order.status === "cancel" ? "You have canceled this order. Thank you for your cooperation." : 
                      order.status === "complete" ? `This order has been completed. If you want to contact your client, please contact via ${order.user.email}` : null
                    }</td>
                    <td className="tw-align-middle">
                      {
                        order.status === "waiting" ? (
                          <div>
                            <Button onClick={() => handleCancel(order.id)} className="!tw-bg-custom-danger hover:!tw-bg-custom-danger-dark !tw-border-0 !tw-mr-4 !tw-mb-4">Cancel</Button>
                            <Button onClick={() => handleApprove(order.id)} className="!tw-bg-custom-success hover:!tw-bg-custom-success-dark !tw-border-0">Approve</Button>
                          </div>
                        ) : 
                        order.status === "cancel" ? <img src="/images/icon-cancel.png" alt={order.status} className="tw-mx-auto" /> : 
                        order.status === "success" ? <Button onClick={() => handleSend(order.id)} className="!tw-bg-custom-success hover:!tw-bg-custom-success-dark !tw-border-0">Send Project</Button> : 
                        order.status === "complete" ? <Button onClick={() => handleView(order.id)} className="!tw-bg-custom-success hover:!tw-bg-custom-success-dark !tw-border-0">View Project</Button> : null
                      }
                    </td>
                  </tr>
                ))
              }
              </tbody>
            </Table>
          ) : (
            <p className="tw-opacity-50">There are no offers to display.</p>
          )
        )}
      </Container>
    </>
  )
}

export default OrderList;