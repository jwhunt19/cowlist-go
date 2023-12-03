import Cow from './Cow';

const Cowlist = ({ updateCow, deleteCow, cows }) => {
  return (
    <div>
      {cows.map((cow) => (
        <Cow updateCow={updateCow} deleteCow={deleteCow} cow={cow} key={cow.Id} />
      ))}
    </div>
  );
};

export default Cowlist;
