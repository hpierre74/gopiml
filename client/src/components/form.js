import React from "react";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import CssBaseline from "@material-ui/core/CssBaseline";
import TextField from "@material-ui/core/TextField";

import LockOutlinedIcon from "@material-ui/icons/AccountCircle";
import Typography from "@material-ui/core/Typography";
import { makeStyles } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

const useStyles = makeStyles(theme => ({
  "@global": {
    body: {
      backgroundColor: theme.palette.common.white
    }
  },
  paper: {
    marginTop: theme.spacing(8),
    display: "flex",
    flexDirection: "column",
    alignItems: "center"
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main
  },
  form: {
    width: "100%", // Fix IE 11 issue.
    marginTop: theme.spacing(1)
  },
  submit: {
    margin: theme.spacing(3, 0, 2)
  }
}));

// To Facebox default face object
const formatImage = (image, name) => ({
  name,
  id: `${name}.jpg`,
  file: image
});

export default function SignIn() {
  const classes = useStyles();
  const [image, setImage] = React.useState(null);
  const [name, setName] = React.useState("Your name");
  const [preview, setPreview] = React.useState([]);

  const handleFileChange = e => {
    const image = e.target.files[0];
    const fileImages = new FileReader();
    fileImages.onload = e => {
      setPreview(e.target.result);
    };

    fileImages.readAsDataURL(image);

    return setImage(formatImage(image, name));
  };

  const handleTextChange = e => setName(e.target.value);

  console.log(preview, image);

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <form
          className={classes.form}
          onSubmit={!!image ? () => alert("ok") : () => alert("not ok")}
          noValidate
        >
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="name"
            label="Your Name"
            name="name"
            autoFocus
            value={name}
            onChange={handleTextChange}
          />

          <input
            accept="image/*"
            id="raised-button-file"
            onChange={handleFileChange}
            multiple
            type="file"
            style={{ display: "none" }}
          />
          <label htmlFor="raised-button-file">
            <Button fullWidth variant="outlined" component="span">
              {`Upload ${image ? image.id : ""}`}
            </Button>
          </label>
          <Button
            fullWidth
            variant="contained"
            color={!!image ? "primary" : "default"}
            className={classes.submit}
          >
            Sign In
          </Button>
        </form>
        <div>
          {preview && (
            <div>
              <img height="100px" width="100px" src={preview} alt="ok" />
            </div>
          )}
        </div>
      </div>
    </Container>
  );
}
