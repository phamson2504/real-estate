import { React, useState, useEffect } from 'react'
import useAuth from '../../hooks/useAuth';
import './Profile.css'

export default function Profile() {
  const { user, authAxios } = useAuth()

  const [currentUser, setCurrentUser] = useState({
    contact: "",
    email: "",
    avatarUrl: "",
    agentName: ""
  });

  console.log(user)

  const [selectedFile, setSelectedFile] = useState(null);

  // const fetchAgent = async () => {
  //   try {
  //     if (user) {
  //       const userResp = await axios.get(`/agent?userId=${user.id}`)
  //       setCurrentUser({ email: user.email, avatarUrl: user.avatar, contact: userResp.data.contact })
  //     }
  //   } catch (error) {
  //     console.error("Error fetching agent:", error);
  //   }
  // }

  useEffect(() => {
    if (user) {
      setCurrentUser({
        contact: user.agent.contact,
        email: user.email,
        avatarUrl: user.avatar,
        agentName: user.agent.name
      })
    }
  }, [user])

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setCurrentUser((prevUser) => ({
      ...prevUser,
      [name]: value,
    }));
  }

  const handleAvatarChange = (event) => {
    const file = event.target.files[0];
    if (file) {
      setSelectedFile(file);

      const reader = new FileReader();
      reader.onload = () => {
        setCurrentUser((prevUser) => ({
          ...prevUser,
          avatarUrl: reader.result,
        }));
      };
      reader.readAsDataURL(file);
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const formData = new FormData();
      formData.append("agentId", user.agent.id)
      formData.append("contact", currentUser.contact);
      formData.append("email", currentUser.email);
      formData.append("agentName", currentUser.agentName);
      if (selectedFile) {
        formData.append("avatar", selectedFile);
      }
      
      await authAxios.post("/authentication/update-profile", formData, {
        headers: { "Content-Type": "multipart/form-data" },
      });
      
      alert("Profile updated successfully!");
    } catch (error) {
      console.error("Error submitting form:", error);
      alert("Failed to update profile.");
    }
  }

  return (
    <div className='pro-file-contain'>
      <h2>Agent Profile</h2>
      <div className='pro-file-content'>
        <form onSubmit={handleSubmit}>
          <div className='pro-file-avatar'>
            <img src={currentUser.avatarUrl||"/user.jpg"} alt='avatar' />
            <h3>{user.username}</h3>
          </div>
          <div className='pro-file-info'>
            <div>
              <label>Email:</label>
              <input
                type="email"
                name="email"
                value={currentUser.email}
                onChange={handleInputChange}
                disabled
              />
            </div>
            <div className='pro-file-info'>
              <div>
                <label>Contact:</label>
                <input
                  type="text"
                  name="contact"
                  value={currentUser.contact}
                  onChange={handleInputChange}
                />
              </div>
            </div>
            <div className='pro-file-info'>
              <div>
                <label>AgentName:</label>
                <input
                  type="text"
                  name="agentName"
                  value={currentUser.agentName}
                  onChange={handleInputChange}
                />
              </div>
            </div>
            <div>
              <label>Avatar:</label>
              <input
                type="file"
                accept="image/*"
                onChange={handleAvatarChange}
              />
            </div>
          </div>
          <button type="submit" className="pro-file-button">Submit</button>
        </form>
      </div>
    </div>

  )
}
