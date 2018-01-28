import React, { Component } from 'react';
 
import {
  StyleSheet,
  Text,
  View,
  TextInput,
  ScrollView
} from 'react-native';
 
import Icon from 'react-native-vector-icons/FontAwesome';

import Container from '../components/Container';
import Button from '../components/Button';

export default class Login extends Component {
  render() {
    return (
<Container>
    <Button 
        styles={{button: styles.transparentButton}}
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