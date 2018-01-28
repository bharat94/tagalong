import React, { Component } from 'react';
import { StackNavigator } from 'react-navigation';
 
import {
  StyleSheet,
  Text,
  View,
} from 'react-native';
 
import Icon from 'react-native-vector-icons/FontAwesome';
import Container from '../components/Container';
import Button from '../components/Button';
import Navigation from './SelectAction'

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
            <Icon name="facebook-official" size={30} color="#3B5699" />
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