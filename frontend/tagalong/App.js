// import React from 'react';
// import LoginPage from './src/pages/LoginPage';
// export default LoginPage

import React, { Component } from 'react';
import { StackNavigator } from 'react-navigation';
 
import {
  StyleSheet,
  Text,
  View,
} from 'react-native';
 
import Container from './src/components/Container';
import Button from './src/components/Button';
import Navigation from './src/pages/SelectAction'

class LoginScreen extends React.Component {
  render() {
    const { navigation } = this.props;
    return (
      <Container>
        <Button 
          styles={{button: styles.transparentButton}}
          onPress={() => navigation.navigate('SelectAction', {})}
        >
          <View style={styles.inline}>
            <Text style={[styles.buttonBlueText, styles.buttonBigText]}>  Connect </Text> 
            <Text style={styles.buttonBlueText}>with Facebook</Text>
          </View>
      </Button>
    </Container>
    );
  }
}

const styles = StyleSheet.create({
transparentButton: {
  marginTop: 600,
    borderColor: '#3B5699',
    borderWidth: 2,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#6684D1'
},
buttonBlueText: {
    fontSize: 20,
    color: '#3B5699'
},
buttonBigText: {
    fontSize: 20,
    fontWeight: 'bold'
},
inline: {
    flexDirection: 'row'
},
});

const LoginPage = StackNavigator({
  Login: {
    screen: LoginScreen,
  },
  SelectAction: {
    screen: Navigation,
    navigationOptions: {
     headerTitle: 'Select Events',
   }
  }
});

export default LoginPage;